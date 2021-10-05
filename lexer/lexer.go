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
	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.EQ,
				Literal: "==",
			}
		} else {
			tok = newToken(token.ASSIGN, l.char)
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{
				Type:    token.NOTEQUAL,
				Literal: "!=",
			}
		} else {
			tok = newToken(token.BANG, l.char)
		}
	case '+':
		tok = newToken(token.PLUS, l.char)
	case '-':
		tok = newToken(token.MINUS, l.char)
	case '*':
		tok = newToken(token.ASTERISK, l.char)
	case '/':
		tok = newToken(token.SLASH, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)
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
	case '[':
		tok = newToken(token.LBRACKET, l.char)
	case ']':
		tok = newToken(token.RBRACKET, l.char)
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Type = token.EOF
	default:
		switch {
		case isLetter(l.char):
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)

			return tok
		case isDigit(l.char):
			tok.Type = token.INT
			tok.Literal = l.readNumber()

			return tok
		default:
			tok.Type = token.ILLEGAL
		}
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

func (l *Lexer) peekChar() byte {
	if l.nextReadPos >= len(l.input) {
		return 0 // EOF
	}

	return l.input[l.nextReadPos]
}

func (l *Lexer) readIdentifier() string {
	identifierStartsAt := l.currentPos
	l.readChar() // doing so we don't need to check current char twice
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[identifierStartsAt:l.currentPos]
}

func (l *Lexer) readNumber() string {
	identifierStartsAt := l.currentPos
	l.readChar() // doing so we don't need to check current char twice
	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[identifierStartsAt:l.currentPos]
}

func (l *Lexer) readString() string {
	stringStartsAt := l.currentPos + 1
	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}

	return l.input[stringStartsAt:l.currentPos]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
