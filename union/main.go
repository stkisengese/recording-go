package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}
	result := ""
	unionStr := os.Args[1] + os.Args[2]

	unionMap := make(map[rune]bool)

	for _, char := range unionStr {
		unionMap[char] = true
	}
	for _, ch := range unionStr {
		if unionMap[ch] {
			result += string(ch)
			// delete(unionMap, ch)
			unionMap[ch] = false
		}
	}
	fmt.Println(result)
	fmt.Println()
}
