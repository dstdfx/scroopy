package evaluator

import (
	"fmt"
	"math"

	"github.com/dstdfx/scroopy/ast"
	"github.com/dstdfx/scroopy/object"
)

var buildInFuncs = map[string]*object.BuildIn{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 1 {
				return newError("wrong number of arguments. got=%d, want=1", lenArgs)
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		}},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 1 {
				return newError("wrong number of arguments. got=%d, want=1", lenArgs)
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arrayObj := args[0].(*object.Array)
			if len(arrayObj.Elements) > 0 {
				return arrayObj.Elements[0]
			}

			return object.NULL
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 1 {
				return newError("wrong number of arguments. got=%d, want=1", lenArgs)
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arrayObj := args[0].(*object.Array)
			arrayLength := len(arrayObj.Elements)
			if arrayLength > 0 {
				return arrayObj.Elements[arrayLength-1]
			}

			return object.NULL
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 1 {
				return newError("wrong number of arguments. got=%d, want=1", lenArgs)
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arrayObj := args[0].(*object.Array)
			arrayLength := len(arrayObj.Elements)
			if arrayLength > 0 {
				newElements := make([]object.Object, arrayLength-1)
				copy(newElements, arrayObj.Elements[1:arrayLength])

				return &object.Array{Elements: newElements}
			}

			return object.NULL
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			lenArgs := len(args)
			if lenArgs != 2 {
				return newError("wrong number of arguments. got=%d, want=2", lenArgs)
			}

			if args[0].Type() != object.ArrayObj {
				return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
			}

			arrayObj := args[0].(*object.Array)
			arrayLength := len(arrayObj.Elements)
			newArray := make([]object.Object, arrayLength+1)
			copy(newArray, arrayObj.Elements)
			newArray[arrayLength] = args[1]

			return &object.Array{Elements: newArray}
		},
	},
}

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
	case *ast.ArrayLiteral:
		elements := evalExpressions(n.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}

		return &object.Array{Elements: elements}
	case *ast.IndexExpression:
		left := Eval(n.Left, env)
		if isError(left) {
			return left
		}

		index := Eval(n.Index, env)
		if isError(index) {
			return index
		}

		return evalIndexExpression(left, index)
	case *ast.HashLiteral:
		return evalHashMapLiteral(n, env)
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
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)

		return unwrapReturnValue(evaluated)
	case *object.BuildIn:
		return fn.Fn(args...)
	default:
		return newError("not a function: %s", fn.Type())
	}
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
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if buildin, ok := buildInFuncs[node.Value]; ok {
		return buildin
	}

	return newError("identifier not found: " + node.Value)
}

func evalInfixExpression(op string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.IntegerObj && right.Type() == object.IntegerObj:
		return evalIntegerInfixExpression(op, left, right)
	case left.Type() == object.StringObj && right.Type() == object.StringObj:
		return evalStringInfixExpression(op, left, right)
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

func evalStringInfixExpression(op string, left, right object.Object) object.Object {
	leftVal := left.(*object.String)
	rightVal := right.(*object.String)

	switch op {
	case "+":
		return &object.String{Value: leftVal.Value + rightVal.Value}
	case "==":
		return &object.Boolean{Value: leftVal.Value == rightVal.Value}
	case "!=":
		return &object.Boolean{Value: leftVal.Value != rightVal.Value}
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
	case "**":
		// TODO: fix me when floats are supported
		return &object.Integer{Value: int64(math.Pow(float64(leftVal.Value), float64(rightVal.Value)))}
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

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.ArrayObj && index.Type() == object.IntegerObj:
		return evalArrayIndexExpression(left, index)
	default:
		// TODO: check index type and write appropriate error msg
		return newError("index operator not supported: %s", left.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObj := array.(*object.Array)
	idx := index.(*object.Integer).Value

	if idx < 0 || idx > int64(len(arrayObj.Elements)-1) {
		return object.NULL
	}

	return arrayObj.Elements[idx]
}

func evalHashMapLiteral(node *ast.HashLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.HashPair)

	for k, v := range node.Pairs {
		key := Eval(k, env)
		if isError(key) {
			return key
		}

		hash, ok := key.(object.Hashable)
		if !ok {
			return newError("unusable as hash key: %s", key.Type())
		}

		value := Eval(v, env)
		if isError(value) {
			return value
		}

		pairs[hash.HashKey()] = object.HashPair{
			Key:   key,
			Value: value,
		}
	}

	return &object.HashMap{ObjType: object.HashObj, Pairs: pairs}
}
