package lexer

import (
	"cc/jsonParser/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `{"name": "kabir", "age": 10}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, `"name"`},
		{token.COLON, ":"},
		{token.STRING, `"kabir"`},
		{token.COMMA, ","},
		{token.KEY, `"age"`},
		{token.COLON, ":"},
		{token.INT, "10"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
