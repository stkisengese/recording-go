package main

import (
	"github.com/01-edu/z01"
)

func ReduceInt(a []int, f func(int, int) int) {
	result := a[0]
	// for i := 1; i < len(a); i++ {
	// 	result = f(result, a[i])
	// }
	for _, val := range a[1:] {
		result = f(result, val)
    }
	PrintStr(Itoa(result))
	
}

func PrintStr(s string) {	
	for _, val := range s {
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
	as := []int{500, 2, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}
