package parser

import (
	"github.com/dstdfx/scroopy/ast"
	"github.com/dstdfx/scroopy/lexer"
	"github.com/dstdfx/scroopy/token"
)

// Parser represents an entity that produces AST.
type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

// New returns new instance of Parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Root {
	// TODO: implement me
	return nil
}
