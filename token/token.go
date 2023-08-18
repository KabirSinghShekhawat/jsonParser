package token

const (
	KEY    = "KEY"
	STRING = "STRING"
	INT    = "INT"

	DOUBLE_QUOTE = `"`
	COLON        = ":"

	LBRACE = "{"
	RBRACE = "}"

	ARRAY_OPEN  = "["
	ARRAY_CLOSE = "]"

	COMMA = ","
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"true":  TRUE,
	"false": FALSE,
}

func LookupKey(key string) TokenType {
	if tok, ok := keywords[key]; ok {
		return tok
	}
	return KEY
}