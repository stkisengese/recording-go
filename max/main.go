package main

import (
	"fmt"
)

func Max(a []int) int {
	max := a[0]
	for _, digit := range a {
		if digit > max {
			max = digit
		}
	}
	return max
}

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}
