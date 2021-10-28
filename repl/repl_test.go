package repl_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dstdfx/scroopy/repl"
)

func TestStart(t *testing.T) {
	input := `let factorial = fn(n) { if (n == 1) { 1 } else { n * factorial(n-1) }}; factorial(5)`
	expected := `>> 120
>> `
	output := bytes.NewBuffer(make([]byte, 0, 32))
	repl.Start(strings.NewReader(input), output)

	if output.String() != expected {
		t.Fail()
	}
}
