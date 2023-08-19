package main

import (
	"fmt"
	"jsonParser/lexer"
	"jsonParser/token"
)

func main() {
	input := `{"name": "kabir", "age": 10, "isActive": true, "friends": ["john", "doe"]}`
	l := lexer.New(input)
	IllegalToken := token.Token{Type: token.ILLEGAL}
	EOFToken := token.Token{Type: token.EOF}
	i := 0
	for tok := l.NextToken(); tok.Type != IllegalToken.Type && tok.Type != EOFToken.Type; tok = l.NextToken() {
		fmt.Printf("%d. (%s) -> (%s) \n", i+1, tok.Type, tok.Literal)
		i += 1
	}
}
