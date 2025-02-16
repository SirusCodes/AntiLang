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
