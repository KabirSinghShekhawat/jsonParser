package main

import (
	"fmt"
	"io"
	"strings"
)

type TokenType string

const (
	ObjectStart   TokenType = "OBJECT_START"
	ObjectEnd     TokenType = "OBJECT_END"
	DoubleQuote   TokenType = "DOUBLE_QUOTE"
	ArrayBegin    TokenType = "ARRAY_BEGIN"
	ArrayEnd      TokenType = "ARRAY_END"
	KvSeparator   TokenType = "KV_SEPARATOR"
	ItemSeparator TokenType = "ITEM_SEPARATOR"
	Boolean       TokenType = "BOOLEAN"
	Number        TokenType = "NUMBER"
	String        TokenType = "STRING"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	tokens []Token
}

func main() {
	jsonReader := strings.NewReader("{\"name\": \"Kabir Singh\"}")
	readBuffer := make([]byte, 1)
	luthor := Lexer{
		tokens: make([]Token, jsonReader.Len()),
	}

	var token strings.Builder

	quoteOpen := false
	for {
		_, err := jsonReader.Read(readBuffer)
		if err == io.EOF {
			break
		}

		token.WriteString(string(readBuffer))

		tokenStr := token.String()
		if tokenStr == "{" {
			luthor.tokens = append(luthor.tokens, Token{Type: ObjectStart, Value: token.String()})
			token.Reset()
		} else if tokenStr == "}" {
			luthor.tokens = append(luthor.tokens, Token{Type: ObjectEnd, Value: token.String()})
			token.Reset()
		} else if string(readBuffer) == "\"" {
			quoteOpen = true
			// HasSuffix is used just for semantic reasons, replacing it with 'Contains' will maken no differnce.
			// Since the token is built char by char, the second double quote will always be a suffix.
			if strings.HasSuffix(tokenStr[1:], "\"") {
				strVal := strings.Split(tokenStr, "\"")[1]
				luthor.tokens = append(luthor.tokens, Token{Type: DoubleQuote, Value: "\""})
				luthor.tokens = append(luthor.tokens, Token{Type: String, Value: strVal})
				luthor.tokens = append(luthor.tokens, Token{Type: DoubleQuote, Value: "\""})
				quoteOpen = false
				token.Reset()
			} else {
				continue
			}
		} else if tokenStr == ":" {
			luthor.tokens = append(luthor.tokens, Token{Type: KvSeparator, Value: token.String()})
			token.Reset()
		} else if string(readBuffer) == " " && !quoteOpen {
			token.Reset()
		}
	}

	for _, token := range luthor.tokens {
		if len(token.Value) != 0 {
			fmt.Println(token)
		}
	}
}
