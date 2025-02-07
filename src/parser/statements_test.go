package parser_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/lexer"
	"github.com/SirusCodes/anti-lang/src/parser"
)

func TestLetStatement(t *testing.T) {
	input := ",5 = five let"

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()

	fmt.Println(program)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(parser.Errors()) != 0 {
		t.Fatalf("parser has %d errors \n %s", len(parser.Errors()), strings.Join(parser.Errors(), "\n"))

	}

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

	// TODO: Test for value of the expression
}
