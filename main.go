package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SirusCodes/anti-lang/src/evaluator"
	"github.com/SirusCodes/anti-lang/src/lexer"
	"github.com/SirusCodes/anti-lang/src/object"
	"github.com/SirusCodes/anti-lang/src/parser"
	"github.com/SirusCodes/anti-lang/src/repl"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	switch command {
	case "repl":
		runREPL()
	case "run":
		if len(os.Args) < 3 {
			printHelp()
			return
		}
		path := os.Args[2]
		runFile(path)
	case "help":
		printHelp()
	default:
		printHelp()
	}
}

func runREPL() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! to AntiLang!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}

func runFile(path string) int {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	l := lexer.New(string(file))
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

func printHelp() {
	fmt.Println("Usage: anti-lang [command] [args]")
	fmt.Println("Commands:")
	fmt.Println("  repl - Start the AntiLang REPL")
	fmt.Println("  run [filename] - Run an AntiLang file")
}
