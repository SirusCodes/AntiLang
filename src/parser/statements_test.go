package parser_test

import (
	"strings"
	"testing"

	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/lexer"
	"github.com/SirusCodes/anti-lang/src/parser"
	"github.com/SirusCodes/anti-lang/src/utils"
)

func TestLetStatement(t *testing.T) {
	input := ",5 + 5 = five let"

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	utils.CheckParserErrors(t, parser)

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

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()

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

func TestBlockStatement(t *testing.T) {
	input := "[,5 + 5 return]"

	lexer := lexer.New(input)
	parser := parser.New(lexer)

	program := parser.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	utils.CheckParserErrors(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d", len(program.Statements))
	}

	stmt := program.Statements[0]

	btt, ok := stmt.(*ast.BlockStatement)

	if !ok {
		t.Fatalf("stmt not *parser.BlockStatement. got=%T", stmt)
	}

	if btt.TokenLiteral() != "[" {
		t.Fatalf("stmt.TokenLiteral not '['. got=%q", btt.TokenLiteral())
	}

	if len(btt.Statements) != 1 {
		t.Fatalf("btt.Statements does not contain 1 statements. got=%d", len(btt.Statements))
	}

	retStmt, ok := btt.Statements[0].(*ast.ReturnStatement)

	if !ok {
		t.Fatalf("stmt not *parser.ReturnStatement. got=%T", btt.Statements[0])
	}

	if retStmt.TokenLiteral() != "return" {
		t.Fatalf("stmt.TokenLiteral not 'return'. got=%q", retStmt.TokenLiteral())
	}

	if retStmt.ReturnValue.String() != "(5 + 5)" {
		t.Fatalf("stmt.ReturnValue.String() not '5 + 5'. got=%q", retStmt.ReturnValue.String())
	}

}
