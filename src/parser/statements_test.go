package parser_test

import (
	"testing"

	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/utils"
)

func TestLetStatement(t *testing.T) {
	input := ",5 + 5 = five let"
	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt := program.Statements[0]
	ltt, ok := stmt.(*ast.LetStatement)

	if !ok {
		t.Fatalf("stmt not *parser.LetStatement. got=%T", stmt)
	}

	if ltt.TokenLiteral() != "let" {
		t.Fatalf("stmt.TokenLiteral not 'let'. got=%q", ltt.TokenLiteral())
	}

	if ltt.Name.Value != "five" {
		t.Fatalf("stmt.Name.Value not 'five'. got=%q", ltt.Name.Value)
	}

	if ltt.Value.String() != "(5 + 5)" {
		t.Fatalf("stmt.Value.String() not '5 + 5'. got=%q", ltt.Value.String())
	}
}

func TestReturnStatement(t *testing.T) {
	input := ",5 + 5 return"
	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt := program.Statements[0]
	rtt, ok := stmt.(*ast.ReturnStatement)

	if !ok {
		t.Fatalf("stmt not *parser.ReturnStatement. got=%T", stmt)
	}

	if rtt.TokenLiteral() != "return" {
		t.Fatalf("stmt.TokenLiteral not 'return'. got=%q", rtt.TokenLiteral())
	}

	if rtt.ReturnValue.String() != "(5 + 5)" {
		t.Fatalf("stmt.ReturnValue.String() not '5 + 5'. got=%q", rtt.ReturnValue.String())
	}
}

func TestMultiStatementInAFunction(t *testing.T) {
	input := `{}abc func [
		,5 = five let
		,55 = five

		,five return
	]`

	program := utils.ParseInput(t, input)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt := program.Statements[0]
	est, ok := stmt.(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("stmt not *ast.ExpressionStatement. got=%T", stmt)
	}

	fn, ok := est.Expression.(*ast.FunctionExpression)
	if !ok {
		t.Fatalf("est.Expression not *ast.FunctionExpression. got=%T", est.Expression)
	}

	if len(fn.Body.Statements) != 3 {
		t.Fatalf("function.Body.Statements does not contain 3 statements. got=%d", len(fn.Body.Statements))
	}

	stmt1, ok := fn.Body.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("stmt1 not *ast.LetStatement. got=%T", fn.Body.Statements[0])
	}

	if stmt1.Name.Value != "five" {
		t.Fatalf("stmt1.Name.Value not 'five'. got=%q", stmt1.Name.Value)
	}

	stmt2, ok := fn.Body.Statements[1].(*ast.ExpressionStatement).Expression.(*ast.AssignExpression)
	if !ok {
		t.Fatalf("stmt2 not *ast.AssignExpression. got=%T", fn.Body.Statements[1])
	}

	if stmt2.String() != "55 = five" {
		t.Fatalf("stmt2.String() not '55 = five'. got=%q", stmt2.String())
	}

	stmt3, ok := fn.Body.Statements[2].(*ast.ReturnStatement)
	if !ok {
		t.Fatalf("stmt3 not *ast.ReturnStatement. got=%T", fn.Body.Statements[2])
	}

	if stmt3.ReturnValue.String() != "five" {
		t.Fatalf("stmt3.ReturnValue.String() not 'five'. got=%q", stmt3.ReturnValue.String())
	}
}
