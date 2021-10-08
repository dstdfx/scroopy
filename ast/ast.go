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

// InfixExpression represents infix expressions, e.x: 5 + 5.
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

/* Basic types */

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

// BooleanLiteral represents string representation of boolean.
type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (b *BooleanLiteral) expressionNode() {}

func (b *BooleanLiteral) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BooleanLiteral) String() string {
	return b.Token.Literal
}

// StringLiteral represents string.
type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}

func (sl *StringLiteral) TokenLiteral() string {
	return sl.Value
}

func (sl *StringLiteral) String() string {
	return sl.Value
}

// IfExpression represents `if-else` expression.
type IfExpression struct {
	Token       token.Token // The `if` token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("if")
	strBuilder.WriteString(ie.Condition.String())
	strBuilder.WriteByte(' ')
	strBuilder.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		strBuilder.WriteString("else ")
		strBuilder.WriteString(ie.Alternative.String())
	}

	return strBuilder.String()
}

// BlockStatement represents block statement within `{}`.
type BlockStatement struct {
	Token      token.Token // The `{` token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}

func (bs *BlockStatement) String() string {
	strBuilder := strings.Builder{}

	for _, s := range bs.Statements {
		strBuilder.WriteString(s.String())
	}

	return strBuilder.String()
}

// FunctionLiteral represents function definition.
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	strBuilder := strings.Builder{}
	params := make([]string, 0)
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	strBuilder.WriteString(fl.TokenLiteral())
	strBuilder.WriteByte('(')
	strBuilder.WriteString(strings.Join(params, ", "))
	strBuilder.WriteString(") ")
	strBuilder.WriteString(fl.Body.String())

	return strBuilder.String()
}

// CallExpression represents function call.
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	strBuilder := strings.Builder{}
	args := make([]string, 0)
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	strBuilder.WriteString(ce.Function.String())
	strBuilder.WriteByte('(')
	strBuilder.WriteString(strings.Join(args, ", "))
	strBuilder.WriteByte(')')

	return strBuilder.String()
}

// ArrayLiteral represents an array containing a list of expressions.
type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (a *ArrayLiteral) expressionNode() {}

func (a *ArrayLiteral) TokenLiteral() string {
	return a.Token.Literal
}

func (a *ArrayLiteral) String() string {
	strBuilder := strings.Builder{}
	strBuilder.WriteByte('[')
	for i, el := range a.Elements {
		strBuilder.WriteString(el.String())
		if i != len(a.Elements)-1 {
			strBuilder.WriteString(", ")
		}
	}
	strBuilder.WriteByte(']')

	return strBuilder.String()
}

// IndexExpression represents index expression: <left-expression>[<index-expression>].
type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	strBuilder := strings.Builder{}

	strBuilder.WriteByte('(')
	strBuilder.WriteString(ie.Left.String())
	strBuilder.WriteByte('[')
	strBuilder.WriteString(ie.Index.String())
	strBuilder.WriteString("])")

	return strBuilder.String()
}

// HashLiteral represents hash literal: { <expression>: <expression> }.
type HashLiteral struct {
	Token token.Token // the '{' token
	Pairs map[Expression]Expression
}

func (hl *HashLiteral) expressionNode() {}

func (hl *HashLiteral) TokenLiteral() string {
	return hl.Token.Literal
}

func (hl *HashLiteral) String() string {
	strBuilder := strings.Builder{}

	strBuilder.WriteByte('{')

	count := 0
	length := len(hl.Pairs)
	for k, v := range hl.Pairs {
		strBuilder.WriteString(k.String())
		strBuilder.WriteString(":")
		strBuilder.WriteString(v.String())

		if count != length-1 {
			strBuilder.WriteString(", ")
		}
		count++
	}

	strBuilder.WriteByte('}')

	return strBuilder.String()
}
