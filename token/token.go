package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers + literals

	IDEN = "IDENT"
	INT  = "INT"

	//OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	//DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"

	LPALEN = "("
	RPALEN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
