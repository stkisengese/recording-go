package main

import (
	"fmt"
	"strings"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return ""
	}
	for _, v := range dec {
		if !(v >= '0' && v <= '9' || v == '.' || v == '-') {
			return dec
		}
	}
	result := ""
	// i := 0
	j := -1

	for _, v := range dec {
		if v == '.' {
			j++
		} else if j == 0 {
			result += string(v)
		} else {
			result += string(v)
		}
	}

	// result := strings.ReplaceAll(dec, ".", "")

	result = strings.TrimPrefix(result, "0")

	return result
}

func main() {
	fmt.Println(NotDecimal("0.123"))
	fmt.Println(NotDecimal("174.2"))
	fmt.Println(NotDecimal("0.1255"))
	fmt.Println(NotDecimal("1.20525856"))
	fmt.Println(NotDecimal("-0.0f00d00"))
	fmt.Println(NotDecimal(""))
	fmt.Println(NotDecimal("-19.525856"))
	fmt.Println(NotDecimal("1952"))
}
