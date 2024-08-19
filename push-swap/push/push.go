// File: push-swap.go

package main

import (
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	elements []int
}

// Main function to handle input and initiate sorting
func main() {
	arg := os.Args[1]
	if len(arg) == 0 {
		return
	}
	args := []string{}
	for _, i := range arg {
		args = append(args, string(arg[i]))
	}

	stackA, err := createStackFromArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	stackB := &Stack{}

	if isSorted(&stackA) {
		return // Already sorted, no operations needed
	}

	operations := sortStacks(&stackA, stackB)

	for _, op := range operations {
		fmt.Println(op)
	}
}

// Function to create stack from command line arguments
func createStackFromArgs(args []string) (Stack, error) {
	stack := Stack{elements: []int{}}
	seen := map[int]bool{}

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return stack, err
		}
		if seen[num] {
			return stack, fmt.Errorf("duplicate number")
		}
		seen[num] = true
		stack.elements = append(stack.elements, num)
	}
	return stack, nil
}

// Function to check if the stack is sorted
func isSorted(stack *Stack) bool {
	for i := 1; i < len(stack.elements); i++ {
		if stack.elements[i-1] > stack.elements[i] {
			return false
		}
	}
	return true
}

// Main sorting logic
func sortStacks(stackA, stackB *Stack) []string {
	operations := []string{}

	for !isSorted(stackA) || len(stackB.elements) > 0 {
		if len(stackA.elements) <= 3 {
			operations = append(operations, sortSmallStack(stackA)...)
		} else {
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
		if stackA.elements[0] > stackA.elements[1] && stackA.elements[1] < stackA.elements[2] && stackA.elements[0] < stackA.elements[2] {
			operations = append(operations, "sa")
		} else if stackA.elements[0] > stackA.elements[1] && stackA.elements[1] > stackA.elements[2] {
			operations = append(operations, "sa", "rra")
		} else if stackA.elements[0] > stackA.elements[1] && stackA.elements[1] < stackA.elements[2] {
			operations = append(operations, "ra")
		} else if stackA.elements[0] < stackA.elements[1] && stackA.elements[1] > stackA.elements[2] && stackA.elements[0] < stackA.elements[2] {
			operations = append(operations, "rra", "sa")
		} else if stackA.elements[0] < stackA.elements[1] && stackA.elements[1] > stackA.elements[2] {
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
