package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/SeaSkyThe/MonkeyInterpreter/lexer"
	"github.com/SeaSkyThe/MonkeyInterpreter/parser"
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
			p := parser.New(l)

			program := p.ParseProgram()

			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				continue
			}

			io.WriteString(out, program.String())
			io.WriteString(out, "\n")

		}
	}
}

const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
