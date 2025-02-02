package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SirusCodes/anti-lang/src/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello, %s! to AntiLang!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
