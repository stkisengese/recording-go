package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}

	word := os.Args[1]
	add := "ay"

	for i, chr := range word {
		for _, vow := range vowels {
			if chr != rune(vow) {
				word = word[:i] + string(chr)
				continue
			} else {
				break
			}
		}
		if len(word) == len(os.Args[1]) {
			PrintStr("No vowels")
			return
		} else {
			word = word + add
		}
	}
	PrintStr(word)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
