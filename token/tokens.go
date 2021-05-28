package token

const (
	// Inner tokens.
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals.
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators.
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords.
	FUNC = "FUNC"
	LET  = "LET"
)

// Type represents token's type.
type Type string

// Token represents a single entity lexed.
type Token struct {
	Type    Type
	Literal string
}
