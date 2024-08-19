package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	i, j := len(slice1)-1, len(slice2)-1
	len1, len2 := len(slice1), len(slice2)

	result := []int{}

	for i >= 0 && j >= 0 {
		if len1 > len2 {
			result = append(result, slice1[i])
			i--
			len1--
		} else if len2 > len1 {
			result = append(result, slice2[j])
			j--
			len2--
		} else {
			result = append(result, slice1[i])
			i--
			len1--
		}
	}
	for i >= 0 {
		result = append(result, slice1[i])
		i--
		len1--
	}
	for j >= 0 {
		result = append(result, slice2[j])
		j--
	}

	return result
}

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}
