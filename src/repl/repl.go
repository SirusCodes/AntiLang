package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SirusCodes/anti-lang/src/evaluator"
	"github.com/SirusCodes/anti-lang/src/lexer"
	"github.com/SirusCodes/anti-lang/src/object"
	"github.com/SirusCodes/anti-lang/src/parser"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator := evaluator.Eval(program, env)

		if evaluator != nil {
			io.WriteString(out, evaluator.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Guess you are not ready for it...\nLet me help you with that with not so useful errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
