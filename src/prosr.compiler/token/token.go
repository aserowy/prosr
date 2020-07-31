package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers
	IDENT = "IDENT"

	// literals
	INT32 = "INT32"

	// operators and punctuation
	ASSIGN    = "="
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	ACTION  = "ACTION"
	HUB     = "HUB"
	MESSAGE = "MESSAGE"
	RETURNS = "RETURNS"
	TO      = "TO"
)

var keywords = map[string]TokenType{
	"action":  ACTION,
	"hub":     HUB,
	"message": MESSAGE,
	"returns": RETURNS,
	"to":      TO,
}

func LookupTokenType(value string) TokenType {
	if tok, ok := keywords[value]; ok {
		return tok
	}
	return IDENT
}
