package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}
	args := os.Args[1]
	fmt.Println(split(args))
}

func split(s string) string {
	slic := ""
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			result += string(s[i])
		} else if s[i] == ' ' && result != "" {
			slic += result + string(' ')
			result = ""
		}
	}
	// if result != "" {
	// 	slic += result
	// }
	res := []rune(result)
	if res[len(res)-1] == ' ' {
	return result[:len(result)-1]
	}
	return slic
}
