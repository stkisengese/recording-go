package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func FoldInt(f func(int, int) int, a []int, n int) {
	result := n
	// for i := 0; i < len(a); i++ {
	// 	result = f(result, a[i])
	// }
	for _, val := range a {
		result = f(result, val)
    }
	PrintStr(strconv.Itoa(result))
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func main() {
	Mul := func(acc int, cur int) int {
		return acc * cur
	}
	Add := func(acc int, cur int) int {
		return acc + cur
	}
	Sub := func(acc int, cur int) int {
		return acc - cur
	}
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	// fmt.Println()
	z01.PrintRune('\n')

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}
