package evaluator

import (
	"github.com/dstdfx/scroopy/ast"
	"github.com/dstdfx/scroopy/object"
)

// TODO: put to object package
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
		return evalRoot(n)
	case *ast.ExpressionStatement:
		return Eval(n.Expression)
	case *ast.BlockStatement:
		return evalBlockStatements(n)
	case *ast.ReturnStatement:
		return &object.ReturnValue{Value: Eval(n.Value)}
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}
	case *ast.BooleanLiteral:
		return boolToBooleanObject(n.Value)
	case *ast.IfExpression:
		return evalIfExpression(n)
	case *ast.PrefixExpression:
		return evalPrefixExpression(n.Operator, Eval(n.Right))
	case *ast.InfixExpression:
		return evalInfixExpression(n.Operator, Eval(n.Left), Eval(n.Right))
	default:
		return NULL
	}
}

func evalInfixExpression(op string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(op, left, right)
	// TODO: get rid of duplicated code
	case op == "==":
		// boolean [infix op] integer
		if left.Type() == object.IntegerObj && right.Type() == object.BooleanObj {
			return evalIntegerBooleanInfixExpression(op, left, right)
		}

		// integer [infix op] boolean
		if left.Type() == object.BooleanObj && right.Type() == object.IntegerObj {
			return evalIntegerBooleanInfixExpression(op, right, left)
		}

		return boolToBooleanObject(left == right)
	case op == "!=":
		// boolean [infix op] integer
		if left.Type() == object.IntegerObj && right.Type() == object.BooleanObj {
			return evalIntegerBooleanInfixExpression(op, left, right)
		}

		// integer [infix op] boolean
		if left.Type() == object.BooleanObj && right.Type() == object.IntegerObj {
			return evalIntegerBooleanInfixExpression(op, right, left)
		}

		return boolToBooleanObject(left != right)
	default:
		return NULL
	}
}

func evalIntegerInfixExpression(op string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer)
	rightVal := right.(*object.Integer)

	switch op {
	case "+":
		return &object.Integer{Value: leftVal.Value + rightVal.Value}
	case "-":
		return &object.Integer{Value: leftVal.Value - rightVal.Value}
	case "/":
		return &object.Integer{Value: leftVal.Value / rightVal.Value}
	case "*":
		return &object.Integer{Value: leftVal.Value * rightVal.Value}
	case ">":
		return boolToBooleanObject(leftVal.Value > rightVal.Value)
	case "<":
		return boolToBooleanObject(leftVal.Value < rightVal.Value)
	case "==":
		return boolToBooleanObject(leftVal.Value == rightVal.Value)
	case "!=":
		return boolToBooleanObject(leftVal.Value != rightVal.Value)
	default:
		// TODO: add error
		return NULL
	}
}

func evalIntegerBooleanInfixExpression(op string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer)
	rightVal := right.(*object.Boolean)
	switch op {
	case "==":
		return boolToBooleanObject((leftVal.Value != 0) == rightVal.Value)
	case "!=":
		return boolToBooleanObject((leftVal.Value != 0) != rightVal.Value)
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
		// TODO: add error
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

func evalRoot(root *ast.Root) object.Object {
	var result object.Object

	for _, s := range root.Statements {
		result = Eval(s)

		if result.Type() == object.ReturnValueObj {
			result = result.(*object.ReturnValue).Value

			break
		}
	}

	return result
}

func evalBlockStatements(block *ast.BlockStatement) object.Object {
	var result object.Object

	for _, s := range block.Statements {
		result = Eval(s)

		if result != nil && result.Type() == object.ReturnValueObj {
			break
		}
	}

	return result
}

func boolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	cond := Eval(ie.Condition)
	if isTruthy(cond) {
		return Eval(ie.Consequence)
	}

	// TODO: maybe just return Eval(ie.Alternative) here, Eval will return NULL anyways
	if ie.Alternative != nil {
		return Eval(ie.Alternative)
	}

	return NULL
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}
