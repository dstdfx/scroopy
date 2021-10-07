package object

import (
	"fmt"
	"strings"

	"github.com/dstdfx/scroopy/ast"
)

const (
	IntegerObj     Type = "INTEGER"
	BooleanObj     Type = "BOOLEAN"
	NullObj        Type = "NULL"
	StringObj      Type = "STRING"
	ReturnValueObj Type = "RETURN_VALUE"
	ErrorObj            = "ERROR"
	FunctionObj         = "FUNCTION"
	BuildInObj          = "BUILDIN"
	ArrayObj            = "ARRAY"
)

var (
	NULL  = &Null{}
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
)

// Type represents object's type.
type Type string

// Object represents an entity being evaluated.
type Object interface {
	Type() Type
	Inspect() string
}

// Integer represents integer type.
type Integer struct {
	Value int64
}

func (i *Integer) Type() Type {
	return IntegerObj
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Boolean represents boolean type.
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type {
	return BooleanObj
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null represents null type.
type Null struct{}

func (n *Null) Type() Type {
	return NullObj
}

func (n *Null) Inspect() string {
	return "null"
}

// String represents string type.
type String struct {
	Value string
}

func (s *String) Type() Type {
	return StringObj
}

func (s *String) Inspect() string {
	return s.Value
}

// ReturnValue represents return statement value.
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() Type {
	return ReturnValueObj
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Error represents an error object.
type Error struct {
	Message string
}

func (e *Error) Type() Type {
	return ErrorObj
}

func (e *Error) Inspect() string {
	return fmt.Sprintf("ERROR: %s", e.Message)
}

// Function represents a function.
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type {
	return FunctionObj
}

func (f *Function) Inspect() string {
	strBuilder := strings.Builder{}

	params := make([]string, 0)
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	strBuilder.WriteString("fn")
	strBuilder.WriteByte('(')
	strBuilder.WriteString(strings.Join(params, ", "))
	strBuilder.WriteString(") {\n")
	strBuilder.WriteString(f.Body.String())
	strBuilder.WriteString("\n}")

	return strBuilder.String()
}

// BuildInFunction represents build-in function definition.
type BuildInFunction func(args ...Object) Object

// BuildIn represents a wrapper around BuildInFunction that implements Object interface.
type BuildIn struct {
	Fn BuildInFunction
}

func (b *BuildIn) Type() Type {
	return BuildInObj
}

func (b *BuildIn) Inspect() string {
	return "builtin function"
}

type Array struct {
	Elements []Object
}

func (ao *Array) Type() Type {
	return ArrayObj
}

func (ao *Array) Inspect() string {
	strBuilder := strings.Builder{}

	strBuilder.WriteByte('[')
	for i, el := range ao.Elements {
		strBuilder.WriteString(el.Inspect())
		if i != len(ao.Elements)-1 {
			strBuilder.WriteString(", ")
		}
	}
	strBuilder.WriteByte(']')

	return strBuilder.String()
}
