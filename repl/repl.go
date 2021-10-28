package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/dstdfx/scroopy/evaluator"
	"github.com/dstdfx/scroopy/lexer"
	"github.com/dstdfx/scroopy/object"
	"github.com/dstdfx/scroopy/parser"
)

// Start runs the main REPL goroutine.
// It reads data from the given io.Reader, parses and evaluates it.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		_, err := io.WriteString(out, ">> ")
		if err != nil {
			handleIOError(err)
		}
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

		evaluated := evaluator.Eval(root, env)
		if evaluated == nil {
			continue
		}

		_, err = io.WriteString(out, evaluated.Inspect()+"\n")
		if err != nil {
			handleIOError(err)
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
