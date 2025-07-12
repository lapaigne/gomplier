package main

import "fmt"

func main() {

	source := "+- */"

	var lexer Lexer
	lexer.init(source)

	token, err := lexer.getToken()
	if err != nil {
		fmt.Println(err)
		fmt.Println("TERMINATING")
		return
	}

	for token.cat != EOF {
		fmt.Println(token.cat)
		token, err = lexer.getToken()

		if err != nil {
			fmt.Println(err)
			fmt.Println("TERMINATING")
			return
		}
	}
}
