package evaluator

import (
	"fmt"

	"github.com/dstdfx/scroopy/ast"
	"github.com/dstdfx/scroopy/object"
)

// Eval function evaluates the given node and returns it's "objective"
// representation.
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch n := node.(type) {
	// Statements
	case *ast.Root:
		return evalRoot(n, env)
	case *ast.ExpressionStatement:
		return Eval(n.Expression, env)
	case *ast.BlockStatement:
		return evalBlockStatements(n, env)
	case *ast.ReturnStatement:
		evaluated := Eval(n.Value, env)
		if isError(evaluated) {
			return evaluated
		}

		return &object.ReturnValue{Value: evaluated}
	case *ast.LetStatement:
		evaluated := Eval(n.Value, env)
		if isError(evaluated) {
			return evaluated
		}
		env.Set(n.Name.Value, evaluated)
	case *ast.FunctionLiteral:
		return &object.Function{
			Parameters: n.Parameters,
			Body:       n.Body,
			Env:        env,
		}

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: n.Value}
	case *ast.BooleanLiteral:
		return boolToBooleanObject(n.Value)
	case *ast.IfExpression:
		return evalIfExpression(n, env)
	case *ast.PrefixExpression:
		evaluated := Eval(n.Right, env)
		if isError(evaluated) {
			return evaluated
		}

		return evalPrefixExpression(n.Operator, evaluated)
	case *ast.InfixExpression:
		rightEvaluated := Eval(n.Right, env)
		if isError(rightEvaluated) {
			return rightEvaluated
		}

		leftEvaluated := Eval(n.Left, env)
		if isError(leftEvaluated) {
			return leftEvaluated
		}

		return evalInfixExpression(n.Operator, leftEvaluated, rightEvaluated)
	case *ast.Identifier:
		return evalIdentifier(n, env)
	case *ast.StringLiteral:
		return &object.String{Value: n.Value}
	case *ast.CallExpression:
		function := Eval(n.Function, env)
		if isError(function) {
			return function
		}
		args := evalExpressions(n.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(function, args)
	}

	return nil
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	return obj != nil && obj.Type() == object.ErrorObj
}

func applyFunction(fn object.Object, args []object.Object) object.Object {
	function, ok := fn.(*object.Function)
	if !ok {
		return newError("not a function: %s", fn.Type())
	}
	extendedEnv := extendFunctionEnv(function, args)
	evaluated := Eval(function.Body, extendedEnv)

	return unwrapReturnValue(evaluated)
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)
	for paramIdx, param := range fn.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	result := make([]object.Object, 0, len(exps))
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}

	return result
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(node.Value)
	if !ok {
		return newError("identifier not found: " + node.Value)
	}

	return val
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
	case left.Type() != right.Type():
		return newError("type mismatch: %s %s %s", left.Type(), op, right.Type())
	default:
		return newError("unknown operator: %s %s %s", left.Type(), op, right.Type())
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
		return newError("unknown operator: %s %s %s", left.Type(), op, right.Type())
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
		return object.NULL
	}
}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return newError("unknown operator: %s%s", op, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case object.TRUE:
		return object.FALSE
	case object.FALSE:
		return object.TRUE
	case object.NULL:
		return object.TRUE
	default:
		return object.FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.IntegerObj {
		return newError("unknown operator: -%s", right.Type())
	}

	return &object.Integer{
		Value: -right.(*object.Integer).Value,
	}
}

func evalRoot(root *ast.Root, env *object.Environment) object.Object {
	var result object.Object

	for _, s := range root.Statements {
		result = Eval(s, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStatements(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, s := range block.Statements {
		result = Eval(s, env)

		if result != nil {
			rt := result.Type()
			if rt == object.ReturnValueObj || rt == object.ErrorObj {
				break
			}
		}
	}

	return result
}

func boolToBooleanObject(input bool) *object.Boolean {
	if input {
		return object.TRUE
	}

	return object.FALSE
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condEvaluated := Eval(ie.Condition, env)
	if isError(condEvaluated) {
		return condEvaluated
	}

	if isTruthy(condEvaluated) {
		return Eval(ie.Consequence, env)
	}

	if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	}

	return object.NULL
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case object.NULL:
		return false
	case object.TRUE:
		return true
	case object.FALSE:
		return false
	default:
		return true
	}
}
