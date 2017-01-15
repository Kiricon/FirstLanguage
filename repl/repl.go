package repl

import (
	"FirstLanguage/lexer"
	"FirstLanguage/token"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := scanner.Scan()
	if !scanner {
		return
	}

	line := sanner.Text()
	l := lexer.new(line)

	for tok := l.NextToke(); token.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("%+v\n", tok)
	}
}
