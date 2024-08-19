package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]


	var result string
	charMap := make(map[rune]bool)

	for _, char := range str2 {
		charMap[char] = true
	}
	for _, char := range str1 {
		if charMap[char] {
			result += string(char)
		}
	}
	fmt.Println(result)
}
