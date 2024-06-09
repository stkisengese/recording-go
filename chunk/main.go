package main

import "fmt"

func Chunk(slice []int, size int) {
	var cuts []int
	if size == 0 {
		fmt.Println()
		return
	}
	var result [][]int

	for len(slice) >= size {
		cuts, slice = slice[:size], slice[size:]
		result = append(result, cuts)
	}
	if len(slice) > 0 {
		result = append(result, slice[:])
	}
	fmt.Println(result)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
