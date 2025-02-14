package utils

import (
	"testing"

	"github.com/SirusCodes/anti-lang/src/ast"
	"github.com/SirusCodes/anti-lang/src/evaluator"
	"github.com/SirusCodes/anti-lang/src/lexer"
	"github.com/SirusCodes/anti-lang/src/object"
	"github.com/SirusCodes/anti-lang/src/parser"
)

func EvalTest(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	return evaluator.Eval(program, env)
}

func ParseInput(t *testing.T, input string) *ast.Program {
	lexer := lexer.New(input)
	parser := parser.New(lexer)
	program := parser.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	checkParserErrors(t, parser)
	return program
}

func checkParserErrors(t *testing.T, p *parser.Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
