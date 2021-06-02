package lexer_test

import (
	"testing"

	"github.com/dstdfx/scroopy/lexer"
	"github.com/dstdfx/scroopy/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);

if (5 < result) {
       return true;
} else {
       return false;
}
`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "five"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "5"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "ten"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.INT, expectedLiteral: "10"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "add"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.FUNC, expectedLiteral: "fn"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "x"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "y"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.IDENT, expectedLiteral: "x"},
		{expectedType: token.PLUS, expectedLiteral: "+"},
		{expectedType: token.IDENT, expectedLiteral: "y"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.RBRACE, expectedLiteral: "}"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.LET, expectedLiteral: "let"},
		{expectedType: token.IDENT, expectedLiteral: "result"},
		{expectedType: token.ASSIGN, expectedLiteral: "="},
		{expectedType: token.IDENT, expectedLiteral: "add"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.IDENT, expectedLiteral: "five"},
		{expectedType: token.COMMA, expectedLiteral: ","},
		{expectedType: token.IDENT, expectedLiteral: "ten"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.IF, expectedLiteral: "if"},
		{expectedType: token.LPAREN, expectedLiteral: "("},
		{expectedType: token.INT, expectedLiteral: "5"},
		{expectedType: token.LT, expectedLiteral: "<"},
		{expectedType: token.IDENT, expectedLiteral: "result"},
		{expectedType: token.RPAREN, expectedLiteral: ")"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.RETURN, expectedLiteral: "return"},
		{expectedType: token.TRUE, expectedLiteral: "true"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.RBRACE, expectedLiteral: "}"},
		{expectedType: token.ELSE, expectedLiteral: "else"},
		{expectedType: token.LBRACE, expectedLiteral: "{"},
		{expectedType: token.RETURN, expectedLiteral: "return"},
		{expectedType: token.FALSE, expectedLiteral: "false"},
		{expectedType: token.SEMICOLON, expectedLiteral: ";"},
		{expectedType: token.RBRACE, expectedLiteral: "}"},
		{expectedType: token.EOF},
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

func TestLexer_NextToken_WithIllegalTokens(t *testing.T) {
	input := `=+-*/(),;{}@!<>==!=`

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
			expectedType:    token.ASTERISK,
			expectedLiteral: "*",
		},
		{
			expectedType:    token.SLASH,
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
			expectedType:    token.BANG,
			expectedLiteral: "!",
		},
		{
			expectedType:    token.LT,
			expectedLiteral: "<",
		},
		{
			expectedType:    token.GT,
			expectedLiteral: ">",
		},
		{
			expectedType:    token.EQ,
			expectedLiteral: "==",
		},
		{
			expectedType:    token.NOTEQUAL,
			expectedLiteral: "!=",
		},
		{
			expectedType:    token.EOF,
			expectedLiteral: "",
		},
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
