package token

const (
	// Inner tokens.
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals.
	IDENT  = "IDENT" // add, foobar, x, y, ...
	INT    = "INT"   // 1343456
	STRING = "STRING"

	// Operators.
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	BANG     = "!"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOTEQUAL = "!="

	// Delimiters.
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords.
	FUNC   = "FUNC"
	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

// Type represents token's type.
type Type string

// Token represents a single entity lexed.
type Token struct {
	Type    Type
	Literal string
}

var keywordsLookup = map[string]Type{
	"fn":     FUNC,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns a type of identifier.
func LookupIdent(ident string) Type {
	if identType, ok := keywordsLookup[ident]; ok {
		return identType
	}

	return IDENT
}
