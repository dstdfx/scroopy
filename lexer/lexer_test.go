package lexer_test

import (
	"testing"

	"github.com/dstdfx/scroopy/lexer"
	"github.com/dstdfx/scroopy/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `=+-*/(),;{}@`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.PLUS,
			expectedLiteral: "+",
		},
		{
			expectedType:    token.MINUS,
			expectedLiteral: "-",
		},
		{
			expectedType:    token.MULTIPLY,
			expectedLiteral: "*",
		},
		{
			expectedType:    token.DIVIDE,
			expectedLiteral: "/",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
		},
		{
			expectedType:    token.ILLEGAL,
			expectedLiteral: "",
		},
		{
			expectedType:    token.EOF,
			expectedLiteral: "",
		},
		//nolint
		// TODO: will be supported later
		//{
		//	expectedType:    token.FUNC,
		//	expectedLiteral: "fn",
		//},
		//{
		//	expectedType:    token.LET,
		//	expectedLiteral: "let",
		//},
	}

	lex := lexer.New(input)

	for idx, test := range tests {
		tok := lex.NextToken()

		if tok.Type != test.expectedType {
			t.Fatalf("test[%d]: expected '%s' token type, but got '%s'",
				idx,
				test.expectedType,
				tok.Type)
		}

		if tok.Literal != test.expectedLiteral {
			t.Fatalf("test[%d]: expected '%s' token literal, but got '%s'",
				idx,
				test.expectedLiteral,
				tok.Literal)
		}
	}
}
