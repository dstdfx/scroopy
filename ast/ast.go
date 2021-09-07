package ast

import (
	"fmt"
	"strings"

	"github.com/dstdfx/scroopy/token"
)

// Node describes a minimal entity of AST.
type Node interface {
	fmt.Stringer
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

func (r *Root) String() string {
	strBuilder := strings.Builder{}

	for _, stmt := range r.Statements {
		strBuilder.WriteString(stmt.String())
	}

	return strBuilder.String()
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

func (i *Identifier) String() string {
	return i.Value
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

func (l *LetStatement) String() string {
	strBuilder := strings.Builder{}

	strBuilder.WriteString(l.TokenLiteral() + " ")
	strBuilder.WriteString(l.Name.String() + " = ")

	if l.Value != nil {
		strBuilder.WriteString(l.Value.String())
	}

	strBuilder.WriteByte(';')

	return strBuilder.String()
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

func (r *ReturnStatement) String() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(r.TokenLiteral() + " ")
	if r.Value != nil {
		strBuilder.WriteString(r.Value.String())
	}
	strBuilder.WriteByte(';')

	return strBuilder.String()
}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) statementNode() {}

// ExpressionStatement represents expression statement.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}

	return ""
}

func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

// IntegerLiteral represents string representation of an integer.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression represents prefix expression e.x: -5, !IsValid(some) and etc.
type PrefixExpression struct {
	Token    token.Token // !, -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	strBuilder := strings.Builder{}

	strBuilder.WriteByte('(')
	strBuilder.WriteString(pe.Operator)
	strBuilder.WriteString(pe.Right.String())
	strBuilder.WriteByte(')')

	return strBuilder.String()
}

// InfixExpression represents infix expressions, e.x: 5 + 5
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {}

func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}

func (oe *InfixExpression) String() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteByte('(')
	strBuilder.WriteString(oe.Left.String())
	strBuilder.WriteString(" " + oe.Operator + " ")
	strBuilder.WriteString(oe.Right.String())
	strBuilder.WriteByte(')')

	return strBuilder.String()
}
