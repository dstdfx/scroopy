package ast

import (
	"github.com/dstdfx/scroopy/token"
)

// Node describes a minimal entity of AST.
type Node interface {
	TokenLiteral() string
}

// Statement describes AST statement e.x: variable and etc.
type Statement interface {
	Node
	statementNode() // TODO: fixme
}

// Expression describes AST expression that produces a value.
type Expression interface {
	Node
	expressionNode() // TODO: fixme
}

// Root represents a root node of program's AST.
type Root struct {
	Statements []Statement
}

func (r *Root) TokenLiteral() string {
	if len(r.Statements) > 0 {
		return r.Statements[0].TokenLiteral()
	}

	return ""
}

// Identifier represents statement's identifier.
type Identifier struct {
	Token token.Token // token.IDENT
	Value string
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) expressionNode() {}

// LetStatement represents `let` statement.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) statementNode() {}

// ReturnStatement represents `return` statement.
type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statementNode() {}
