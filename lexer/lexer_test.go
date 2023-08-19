package lexer

import (
	"jsonParser/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `{"name": "kabir", "age": 10, "isActive": true, "friends": ["john", "doe"]}`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, `"name"`},
		{token.COLON, ":"},
		{token.STRING, `"kabir"`},
		{token.COMMA, ","},
		{token.STRING, `"age"`},
		{token.COLON, ":"},
		{token.INT, "10"},
		{token.COMMA, ","},
		{token.STRING, `"isActive"`},
		{token.COLON, ":"},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.STRING, `"friends"`},
		{token.COLON, ":"},
		{token.ARRAY_OPEN, "["},
		{token.STRING, `"john"`},
		{token.COMMA, ","},
		{token.STRING, `"doe"`},
		{token.ARRAY_CLOSE, "]"},
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
