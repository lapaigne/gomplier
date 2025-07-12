package main

import "fmt"

func main() {

	source := "+-*/"

	var lexer Lexer
	lexer.init(source)

	token := lexer.getToken()

	fmt.Println(lexer.char)

	for token.cat != EOF {
		fmt.Println(token.cat)
		token = lexer.getToken()
	}
}
