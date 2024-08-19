package repl

import (
	"bufio"
	"fmt"
	"github.com/andreyrcdias/monkey/lexer"
	"github.com/andreyrcdias/monkey/token"
	"io"
)

const MONKEY_LOGO = `
               o
               | /
o-O-o o-o o-o  OO   o-o o  o
| | | | | |  | | \  |-' |  |
o o o o-o o  o o  o o-o o--O
                           |
                        o--o
`
const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	// TODO: how to proper version a golang package
	version := "0.0.1"

	fmt.Print(MONKEY_LOGO)
	fmt.Printf("Monkey Language %s\n", version)

	fmt.Printf("Type \"help\", \"credits\" for more information.\n")
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		switch line {
		case "help":
			fmt.Printf("TODO")
		case "credits":
			fmt.Printf(" Thanks Thorsten Ball for this Amazin book. \n")
		default:
			l := lexer.New(line)
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				fmt.Printf("%+v\n", tok)
			}
		}

	}
}
