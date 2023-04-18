package repl

import (
	"Monkey/lexer"
	"Monkey/token"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextTokne(); tok.Type != token.EOF; tok = l.NextTokne() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
