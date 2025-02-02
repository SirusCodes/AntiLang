package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/SirusCodes/anti-lang/src/lexer"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != lexer.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
