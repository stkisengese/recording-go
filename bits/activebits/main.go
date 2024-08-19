package main

import "fmt"

// func ActiveBits(n int) int {
// 	count := 0
//     for n != 0 {
//         n = n & (n - 1) // clear the right most bit
//         count++
//     }
//     return count
// }
// func ActiveBits(n int) int {
// 	count := 0
// 	for n > 0 {
// 		count += n & 1 // check the least significant bit
// 		n >>= 1 // shift right by 1 bit (dicard the least significant bit)
// 	}
// 	return count
// }

func ActiveBits(n int) int {
	if n == 0	{
		return 0
	}
	//return ActiveBits(n >> 1) + (n & 1)
	return 1 + ActiveBits(n & (n - 1))
}

func main() {
	fmt.Println(ActiveBits(47))
}
