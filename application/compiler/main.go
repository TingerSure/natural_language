package main

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler"
	"os"
)

func main() {
	source, err := os.Open("test.nl")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	comp := compiler.NewComplier()
	lexer := comp.GetLexer()
	tokens, err := lexer.Read(source)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	tokens = comp.TokenTrim(tokens)
	for _, token := range tokens {
		fmt.Printf("%v | %v\n", token.Type(), token.Value())
	}
}
