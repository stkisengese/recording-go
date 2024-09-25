// File: push-swap.go

package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Stack struct holding slice of integers
type Stack struct {
	elements []int
}

// Push method for Stack
func (s *Stack) Push(value int) {
	s.elements = append(s.elements, value)
}

// Pop method for Stack
func (s *Stack) Pop() int {
	n := len(s.elements) - 1
	value := s.elements[n]
	s.elements = s.elements[:n]
	return value
}

// // Swap method for Stack
// func (s *Stack) Swap() {
// 	if len(s.elements) >= 2 {
// 		slices.Swap(s.elements, 0, 1)
// 	}
// }

// // Reverse method for Stack using slices.Reverse
// func (s *Stack) Reverse() {
// 	slices.Reverse(s.elements)
// }

// // Rotate (Shift Up) method for Stack
// func (s *Stack) Rotate() {
// 	if len(s.elements) > 0 {
// 		slices.Rotate(s.elements, 1)
// 	}
// }

// // Reverse Rotate (Shift Down) method for Stack
// func (s *Stack) ReverseRotate() {
// 	if len(s.elements) > 0 {
// 		slices.Rotate(s.elements, -1)
// 	}
// }

// IsSorted method to check if the Stack is sorted
func (s *Stack) IsSorted() bool {
	return slices.IsSorted(s.elements)
}

// Main function to handle input and initiate sorting
func main() {
	arg := os.Args[1]
	if len(arg) == 0 {
		return
	}
	input := os.Args[1]
	fmt.Println(input)

	stackA, err := createStackFromArgs(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	fmt.Println("StackA created: ", stackA)
	stackB := &Stack{}

	if stackA.IsSorted() {
		fmt.Println("sorted stack")
		return // Already sorted, no operations needed
	}

	operations := sortStacks(stackA, stackB)

	for _, op := range operations {
		fmt.Println(op)
	}
}

// Function to create stack from command line arguments
func createStackFromArgs(input string) (*Stack, error) {
	numbers := strings.Fields(input)
	stack := &Stack{elements: make([]int, 0, len(numbers))}
	seen := make(map[int]bool)

	for _, numStr := range numbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %v", numStr)
		}
		if seen[num] {
			return nil, fmt.Errorf("duplicate number: %d", num)
		}
		seen[num] = true
		stack.elements = append(stack.elements, num)
	}
	fmt.Printf("Input convertion completed successfully")
	fmt.Printf("new elements: %v\n", stack.elements)
	fmt.Printf("new stack: %v\n", stack)
	return stack, nil
}

// Main sorting logic
func sortStacks(stackA, stackB *Stack) []string {
	operations := []string{}

	if !stackA.IsSorted() || len(stackB.elements) > 0 {
		// fmt.Println("sorting validation...")
		if len(stackA.elements) <= 3 {
			// fmt.Println("Elements less than 3\n\n\n")
			operations = append(operations, sortSmallStack(stackA)...)
		} else if slices.Min(stackA.elements) == stackA.elements[len(stackA.elements)-1] {
			operations = append(operations, "rra", "pb")
			stackA.elements = stackA.elements[:len(stackA.elements)-1]
			stackB.elements = append(stackB.elements, stackA.elements[len(stackA.elements)-1])
			if len(stackA.elements) <= 3 {
				// sortSmallStack(stackA)
				operations = append(operations, sortSmallStack(stackA)...)
			}

		} else {
			fmt.Println("Elements more than 3")
			operations = append(operations, greedyMove(stackA, stackB)...)
		}
	}

	operations = append(operations, mergeBack(stackA, stackB)...)
	// return optimizeOperations(operations)
	return operations
}

// Function to handle greedy move based on the calculated greediness
func greedyMove(stackA, stackB *Stack) []string {
	operations := []string{}
	minGreediness := int(^uint(0) >> 1) // Max int value
	bestIndex := -1

	// Find the element with the lowest greediness
	for i := range stackA.elements {
		greediness := calculateGreediness(stackA, i)
		if greediness < minGreediness {
			minGreediness = greediness
			bestIndex = i
		}
	}

	// Perform the best move based on greediness
	if bestIndex != -1 {
		if bestIndex == 0 {
			operations = append(operations, "pb")
			stackB.elements = append([]int{stackA.elements[0]}, stackB.elements...)
			stackA.elements = stackA.elements[1:]
		} else {
			rotateCount := min(bestIndex, len(stackA.elements)-bestIndex)
			if bestIndex <= len(stackA.elements)/2 {
				for i := 0; i < rotateCount; i++ {
					operations = append(operations, "ra")
					stackA.elements = append(stackA.elements[1:], stackA.elements[0])
				}
			} else {
				for i := 0; i < rotateCount; i++ {
					operations = append(operations, "rra")
					stackA.elements = append([]int{stackA.elements[len(stackA.elements)-1]}, stackA.elements[:len(stackA.elements)-1]...)
				}
			}
			operations = append(operations, "pb")
			stackB.elements = append([]int{stackA.elements[0]}, stackB.elements...)
			stackA.elements = stackA.elements[1:]
		}
	}

	return operations
}

// Merging stackB back into stackA
func mergeBack(stackA, stackB *Stack) []string {
	operations := []string{}
	for len(stackB.elements) > 0 {
		operations = append(operations, "pa")
		stackA.elements = append([]int{stackB.elements[0]}, stackA.elements...)
		stackB.elements = stackB.elements[1:]
	}
	return operations
}

// Sorting logic for small stacks (3 or fewer elements)
func sortSmallStack(stackA *Stack) []string {
	operations := []string{}

	if len(stackA.elements) == 2 {
		if stackA.elements[0] > stackA.elements[1] {
			operations = append(operations, "sa")
		}
	} else if len(stackA.elements) == 3 {
		// Sorting logic for exactly 3 elements
		minElem := slices.Min(stackA.elements)
		maxElem := slices.Max(stackA.elements)

		switch {
		case minElem == stackA.elements[0]:
			operations = append(operations, "rra", "sa")
		case maxElem == stackA.elements[2]:
			operations = append(operations, "sa")
		case minElem == stackA.elements[1]:
			operations = append(operations, "ra")
		case maxElem == stackA.elements[0]:
			operations = append(operations, "sa", "rra")
		case minElem == stackA.elements[2]:
			operations = append(operations, "rra")
		}
	}
	return operations
}

// Function to calculate the greediness of moving an element to stack B
func calculateGreediness(stack *Stack, index int) int {
	greedinessValue := 0

	element := stack.elements[index]

	// Factor 1: Distance to the correct position in stack B
	greedinessValue += calculateDistanceToPosition(stack, index)

	// Factor 2: Number of elements that could be correctly placed after this move
	greedinessValue -= calculateFutureImpact(stack, element)

	// Factor 3: Other relevant factors
	// greedinessValue += calculateOtherFactors(stack, element)

	return greedinessValue
}

// Helper function to calculate distance to the correct position in stack B
func calculateDistanceToPosition(stack *Stack, index int) int {
	element := stack.elements[index]
	distance := 0

	for i := 0; i < len(stack.elements); i++ {
		if stack.elements[i] < element {
			distance++
		}
	}

	return min(distance, len(stack.elements)-distance)
}

// Helper function to calculate the future impact of the move
func calculateFutureImpact(stack *Stack, element int) int {
	impact := 0

	for _, val := range stack.elements {
		if val > element {
			impact++
		}
	}

	return impact
}

// Helper function to calculate other factors
// func calculateOtherFactors(stack *Stack, element int) int {
// 	// Placeholder logic for additional factors
// 	return 0
// }

// Utility function to find the minimum of two values
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// // Optimize operations to remove redundancies
// func optimizeOperations(operations []string) []string {
// 	// Placeholder: Just returning the original operations
// 	return operations
// }
