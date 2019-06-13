package evaluator

import (
	"byaigo/ast"
	"byaigo/object"
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// Statments
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statment := range stmts {
		result = Eval(statment)
	}
	return result
}
