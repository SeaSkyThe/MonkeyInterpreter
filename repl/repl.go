package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/SeaSkyThe/MonkeyInterpreter/lexer"
	"github.com/SeaSkyThe/MonkeyInterpreter/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		if strings.ToLower(line) == "clear" {
			fmt.Println("\033[H\033[2J")
		} else {
			l := lexer.New(line)

			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}

		}
	}
}
