package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Params := os.Args[1:]
	for _, arguments := range Params {
		for _, value := range arguments {
			z01.PrintRune(value)
		}
		z01.PrintRune('\n')
	}
}
