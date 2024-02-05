package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/SeaSkyThe/MonkeyInterpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Type any command you want\n")

	repl.Start(os.Stdin, os.Stdout)
}
