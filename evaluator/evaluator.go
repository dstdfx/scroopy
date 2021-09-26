package evaluator

import (
	"github.com/dstdfx/scroopy/ast"
	"github.com/dstdfx/scroopy/object"
)

// Eval function evaluates the given node and returns it's "objective"
// representation.
func Eval(node ast.Node) object.Object {
	switch n := node.(type) {
	// Statements
	case *ast.Root:
		return evalStatements(n.Statements)
	case *ast.ExpressionStatement:
		return Eval(n.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}
	case *ast.BooleanLiteral:
		return &object.Boolean{Value: n.Value}
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, s := range statements {
		result = Eval(s)
	}

	return result
}
