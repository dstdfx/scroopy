package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dstdfx/scroopy/lexer"
	"github.com/dstdfx/scroopy/token"
)

// Start runs the main REPL goroutine.
// It reads data from the given io.Reader and prints parsed tokens (just for now).
func Start(in io.Reader, _ io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		tok := l.NextToken()
		for ; tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
