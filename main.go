package main

import (
	"fmt"
	"io"
	"strings"
)

type Token struct {
	Name  string
	Value string // rune or string?
	Count int
}

type Lexer[T any] struct {
	tokens []T
}

func main() {
	leftBracket := Token{
		Name:  "leftBracket",
		Value: "{",
		Count: 0,
	}

	lexLuthor := Lexer[Token]{
		tokens: make([]Token, 2),
	}

	jsonStr := strings.NewReader("{\"name\": \"Kabir\"}")
	token := make([]byte, 1)
	for {
		_, err := jsonStr.Read(token)
		if err == io.EOF {
			break
		}
		fmt.Printf("token = %s\n", string(token))
		switch string(token) {
		case leftBracket.Value:
			leftBracket.Count += 1
			lexLuthor.tokens = append(lexLuthor.tokens, leftBracket)
		}
	}
	fmt.Printf("\ntokens = %v\n", lexLuthor.tokens)
}
