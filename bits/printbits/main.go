package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Please provide a single integer argument.")
		os.Exit(1)
	}
	arg := os.Args[1]
	num, err := strconv.Atoi(arg)
	// num, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		fmt.Printf("00000000\n")
		os.Exit(0)
	}
	// fmt.Printf("%08b\n", num)
	fmt.Println(manualBinary(num))
}

func manualBinary(num int) string {
	if num == 0 {
		return "0"
	}
	result := ""
	for num > 0 {
		// result = strconv.Itoa(num%2) + result
		result = string((num%2)+'0') + result
		num /= 2
	}
	return result
}
