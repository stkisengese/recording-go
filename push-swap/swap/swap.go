// File: checker.go

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	if stackA.isSorted() && len(stackB.elements) == 0 {
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
			stackA.swap()
		case "sb":
			stackB.swap()
		case "ss":
			stackA.swap()
			stackB.swap()
		case "pa":
			if len(stackB.elements) > 0 {
				stackA.push(stackB.pop())
			}
		case "pb":
			if len(stackA.elements) > 0 {
				stackB.push(stackA.pop())
			}
		case "ra":
			stackA.rotate()
		case "rb":
			stackB.rotate()
		case "rr":
			stackA.rotate()
			stackB.rotate()
		case "rra":
			stackA.reverseRotate()
		case "rrb":
			stackB.reverseRotate()
		case "rrr":
			stackA.reverseRotate()
			stackB.reverseRotate()
		}
	}
}

// Swap the first two elements of the stack
func (s *Stack) swap() {
	if len(s.elements) >= 2 {
		s.elements[0], s.elements[1] = s.elements[1], s.elements[0]
	}
}

// Push an element onto the stack
func (s *Stack) push(element int) {
	s.elements = append([]int{element}, s.elements...)
}

// Pop an element from the stack
func (s *Stack) pop() int {
	element := s.elements[0]
	s.elements = s.elements[1:]
	return element
}

// Rotate the stack (shift up)
func (s *Stack) rotate() {
	if len(s.elements) > 0 {
		s.elements = append(s.elements[1:], s.elements[0])
	}
}

// Reverse rotate the stack (shift down)
func (s *Stack) reverseRotate() {
	if len(s.elements) > 0 {
		// stack.elements = append([]int{stack.elements[len(stack.elements)-1]}, stack.elements[:len(stack.elements)-1]...)
		slices.Reverse(s.elements)
	}
}

// Check if the stack is sorted in ascending order
func (s *Stack) isSorted() bool {
	return slices.IsSorted(s.elements)
}
