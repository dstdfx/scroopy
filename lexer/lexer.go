package lexer

import "github.com/dstdfx/scroopy/token"

// Lexer takes source code as an input and tokenizes it.
type Lexer struct {
	input       string // input data to tokenize
	currentPos  int    // current position in input, index of the char
	nextReadPos int    // current next reading position after the currentPos
	char        byte   // current read char
}

// New returns new instance of Lexer.
func New(src string) *Lexer {
	l := &Lexer{input: src}
	l.readChar() // in order to initialize lexer fields

	return l
}

// NextToken method returns next token in the input.
// If there's no more tokens left - token with token.EOF type is returned.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN, l.char)
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '*':
		tok = newToken(token.MULTIPLY, l.char)
	case '/':
		tok = newToken(token.DIVIDE, l.char)
	case '(':
		tok = newToken(token.LPAREN, l.char)
	case ')':
		tok = newToken(token.RPAREN, l.char)
	case ',':
		tok = newToken(token.COMMA, l.char)
	case ';':
		tok = newToken(token.SEMICOLON, l.char)
	case '{':
		tok = newToken(token.LBRACE, l.char)
	case '}':
		tok = newToken(token.RBRACE, l.char)
	case 0:
		tok.Type = token.EOF
	default:
		tok.Type = token.ILLEGAL
	}
	l.readChar()

	return tok
}

func newToken(tokenType token.Type, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}

func (l *Lexer) readChar() {
	if l.nextReadPos >= len(l.input) {
		l.char = 0 // EOF
	} else {
		l.char = l.input[l.nextReadPos]
	}
	l.currentPos = l.nextReadPos
	l.nextReadPos++
}
