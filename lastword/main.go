package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	if os.Args[1] == "" {
		return
	}
	text := os.Args[1]
	if len(text) == 1 && string(text[0]) == " " {
		return
	}
	lastWord := ""

	for i := len(text) - 1; i >= 0; i-- {
		if string(text[i]) != " " {
			lastWord = string(text[i]) + lastWord
		} else if lastWord != "" {
			break
		}
	}
	fmt.Println(lastWord)
}
