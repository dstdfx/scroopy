package parser

import (
	"fmt"

	"github.com/dstdfx/scroopy/ast"
	"github.com/dstdfx/scroopy/lexer"
	"github.com/dstdfx/scroopy/token"
)

var ErrExpectedNextTokenFmt = "expected next token to be '%s', got '%s' instead"

// Parser represents an entity that produces AST.
type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
	errors       []string
}

// New returns new instance of Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: make([]string, 0)}

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Errors method returns a slice of encountered errors.
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.Type) {
	p.errors = append(p.errors, fmt.Sprintf(ErrExpectedNextTokenFmt, t, p.peekToken.Type))
}

// ParseProgram method parses the program and builds AST.
func (p *Parser) ParseProgram() *ast.Root {
	root := &ast.Root{}
	root.Statements = make([]ast.Statement, 0)

	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			root.Statements = append(root.Statements, stmt)
		}
		p.nextToken()
	}

	return root
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{
		Token: p.currentToken,
	}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: add proper expression parsing
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekToken.Type == t {
		p.nextToken()

		return true
	}

	p.peekError(t)

	return false
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}
	p.nextToken()

	// TODO: add proper expression parsing
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}
