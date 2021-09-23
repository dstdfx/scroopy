package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dstdfx/scroopy/lexer"
	"github.com/dstdfx/scroopy/parser"
	"github.com/dstdfx/scroopy/token"
)

// Start runs the main REPL goroutine.
// It reads data from the given io.Reader and prints parsed AST.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		root := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())

			continue
		}

		_, err := io.WriteString(out, root.String()+"\n")
		if err != nil {
			handleIOError(err)
		}

		tok := l.NextToken()
		for ; tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, err := io.WriteString(out, "\t"+msg+"\n")
		if err != nil {
			handleIOError(err)
		}
	}
}

func handleIOError(err error) {
	_, _ = fmt.Fprintf(os.Stderr, "failed to print program with error: %s\n", err)
	os.Exit(1)
}
