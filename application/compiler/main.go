package main

import (
	"fmt"
	"github.com/TingerSure/natural_language/core/compiler"
	"os"
)

func main() {

	comp := compiler.NewComplier()

	fmt.Println(comp.GetGrammar().GetTable().ToString())
	fmt.Println()

	source, err := os.Open("test2.nl")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	phrase, err := comp.Read(source)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("```")
	fmt.Println(phrase.ToString(""))
	fmt.Println("```")

}
