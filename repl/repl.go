package repl

import (
	"Monkey/evaluator"
	"Monkey/lexer"
	"Monkey/parser"
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
		lexer := lexer.New(line)
		parser := parser.New(lexer)
		program := parser.ParseProgram()

		obj := evaluator.Eval(program)

		fmt.Println(program.String())
		fmt.Println(obj)
	}
}
