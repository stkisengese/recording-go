package main

import (
	"fmt"
)

func FifthAndSkip(str string) string {
	var runes []rune
	count := 0
	for _, r := range str {
		if 33 <= r && r <= 126 {
			if count < 5 {
				runes = append(runes, r)
				count++
			} else {
				runes = append(runes, ' ')
				count = 0
			}
		}
	}
	if len(runes) < 5 {
		return "Invalid Input\n"
	}

	return string(runes) + "\n"
}

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}
