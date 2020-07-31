package repl

import (
	"bufio"
	"fmt"
	"io"

	"prosr.compiler/lexer"
	"prosr.compiler/token"
)

// Start streams file content to the lexer
func Start(in io.Reader, out io.Writer) {
	const PROMPT = ">> "

	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
