//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/SirusCodes/anti-lang/src/evaluator"
	"github.com/SirusCodes/anti-lang/src/lexer"
	"github.com/SirusCodes/anti-lang/src/object"
	"github.com/SirusCodes/anti-lang/src/parser"
)

func main() {
	js.Global().Set("execute", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}

		return execute(args[0].String())
	}))

	select {}
}

func execute(input string) int {
	l := lexer.New(input)
	p := parser.New(l)
	ast := p.ParseProgram()

	if len(p.Errors()) != 0 {
		for _, msg := range p.Errors() {
			fmt.Println(msg)
		}
		return 1
	}

	env := object.NewEnvironment()
	resp := evaluator.Eval(ast, env)

	if resp != nil && resp.Type() == object.ERROR_OBJ {
		fmt.Println(resp.Inspect())
		return 1
	}

	return 0
}
