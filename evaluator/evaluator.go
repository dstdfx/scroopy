package evaluator

import (
	"github.com/dstdfx/scroopy/ast"
	"github.com/dstdfx/scroopy/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
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
	case *ast.PrefixExpression:
		return evalPrefixExpression(n.Operator, Eval(n.Right))
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}
	case *ast.BooleanLiteral:
		return boolToBooleanObject(n.Value)
	default:
		return NULL
	}
}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.IntegerObj {
		return NULL
	}

	return &object.Integer{
		Value: -right.(*object.Integer).Value,
	}
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, s := range statements {
		result = Eval(s)
	}

	return result
}

func boolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}
