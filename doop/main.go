package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	val1 := Atoi(os.Args[1])
	val2 := Atoi(os.Args[3])

	var result int
	switch os.Args[2] {
	case "+":
		result = (val1 + val2)
	case "-":
		result = (val1 - val2)
	case "*":
		result = (val1 * val2)
	case "/":
		if val2 == 0 {
			os.Stdout.WriteString("No division by 0")
			return
		}
		result = (val1 / val2)
	case "%":
		if val2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = (val1 % val2)
	default:
		return
	}
	if (val1 < 0 && val2 < 0 && val1 < (-1<<63)-val2) || (val1 > 0 && val2 > 0 && val1 > (1<<63-1)-val2) {
		return
	}

	os.Stdout.WriteString(Itoa(result) + "\n")
}

func Itoa(num int) string {
	numStr := ""
	if num == 0 {
		return "0"
	}
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}
	for num > 0 {
		digit := num % 10
		numStr = string('0'+digit) + numStr
		num /= 10
	}
	if isNegative {
		numStr = "-" + numStr
	}

	return numStr
}

func Atoi(numStr string) int {
	if numStr == "0" {
		return 0
	}
	isNegative := false
	if numStr[0] == '-' {
		isNegative = true
		numStr = numStr[1:]
	} else if numStr[0] == '+' {
		numStr = numStr[1:]
	}
	result := 0
	if len(numStr) > 0 {
		for _, num := range numStr {
			if num >= '0' && num <= '9' {
				digit := int(num - '0')
				result = result*10 + digit
			} else {
				os.Exit(0)
			}
		}
	}
	if isNegative {
		result = -1 * result
	}
	return result
}
