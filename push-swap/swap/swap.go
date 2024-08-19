// File: checker.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack struct {
	elements []int
}

// Main function to handle input, operations, and validation
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	stackA, err := createStackFromArgs(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}
	stackB := &Stack{}

	operations, err := readOperationsFromStdin()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	executeOperations(stackA, stackB, operations)

	if isSorted(stackA) && len(stackB.elements) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

// Function to create stack from command line arguments
func createStackFromArgs(args []string) (*Stack, error) {
	stack := &Stack{elements: []int{}}
	seen := map[int]bool{}

	for _, arg := range args {
		num, err := strconv.Atoi(arg)
		if err != nil {
			return nil, err
		}
		if seen[num] {
			return nil, fmt.Errorf("duplicate number")
		}
		seen[num] = true
		stack.elements = append(stack.elements, num)
	}
	return stack, nil
}

// Function to read operations from stdin
func readOperationsFromStdin() ([]string, error) {
	operations := []string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if !isValidOperation(line) {
			return nil, fmt.Errorf("invalid operation")
		}
		operations = append(operations, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}

// Function to validate if the operation is one of the allowed operations
func isValidOperation(op string) bool {
	validOps := map[string]bool{
		"sa": true, "sb": true, "ss": true,
		"pa": true, "pb": true,
		"ra": true, "rb": true, "rr": true,
		"rra": true, "rrb": true, "rrr": true,
	}
	return validOps[op]
}

// Function to execute the operations on the stacks
func executeOperations(stackA, stackB *Stack, operations []string) {
	for _, op := range operations {
		switch op {
		case "sa":
			swap(stackA)
		case "sb":
			swap(stackB)
		case "ss":
			swap(stackA)
			swap(stackB)
		case "pa":
			if len(stackB.elements) > 0 {
				push(stackA, pop(stackB))
			}
		case "pb":
			if len(stackA.elements) > 0 {
				push(stackB, pop(stackA))
			}
		case "ra":
			rotate(stackA)
		case "rb":
			rotate(stackB)
		case "rr":
			rotate(stackA)
			rotate(stackB)
		case "rra":
			reverseRotate(stackA)
		case "rrb":
			reverseRotate(stackB)
		case "rrr":
			reverseRotate(stackA)
			reverseRotate(stackB)
		}
	}
}

// Swap the first two elements of the stack
func swap(stack *Stack) {
	if len(stack.elements) >= 2 {
		stack.elements[0], stack.elements[1] = stack.elements[1], stack.elements[0]
	}
}

// Push an element onto the stack
func push(stack *Stack, element int) {
	stack.elements = append([]int{element}, stack.elements...)
}

// Pop an element from the stack
func pop(stack *Stack) int {
	element := stack.elements[0]
	stack.elements = stack.elements[1:]
	return element
}

// Rotate the stack (shift up)
func rotate(stack *Stack) {
	if len(stack.elements) > 0 {
		stack.elements = append(stack.elements[1:], stack.elements[0])
	}
}

// Reverse rotate the stack (shift down)
func reverseRotate(stack *Stack) {
	if len(stack.elements) > 0 {
		stack.elements = append([]int{stack.elements[len(stack.elements)-1]}, stack.elements[:len(stack.elements)-1]...)
	}
}

// Check if the stack is sorted in ascending order
func isSorted(stack *Stack) bool {
	for i := 1; i < len(stack.elements); i++ {
		if stack.elements[i-1] > stack.elements[i] {
			return false
		}
	}
	return true
}
