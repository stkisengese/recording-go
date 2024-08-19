package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	result := ""
	charMap := make(map[rune]bool)

	for _, char := range str2 {
		charMap[char] = true
	}
	for _, char1 := range str1 {
		if  charMap[char1] {
			result += string(char1)
			
			delete(charMap, char1)
		}
	}
	PrintStr(result)
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
