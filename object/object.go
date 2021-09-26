package object

import "fmt"

const (
	IntegerObj Type = "INTEGER"
	BooleanObj Type = "BOOLEAN"
	NullObj    Type = "NULL"
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
