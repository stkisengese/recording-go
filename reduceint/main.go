package main

import (
	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	result := f(a[0], a[1])
	resultStr := Itoa(result)
	for _, val := range resultStr {
		z01.PrintRune(val)
	}
	z01.PrintRune('\n')
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

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
