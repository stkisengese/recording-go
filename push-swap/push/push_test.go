// File: push-swap.go

package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_createStackFromArgs(t *testing.T) {
	args := []string{"3", "2", "1"}
	expected := Stack{elements: []int{1, 2, 3}}

	stack, err := createStackFromArgs(args)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !equalStacks(stack, expected) {
		t.Errorf("Expected stack to be %v, got %v", expected, stack)
	}
}
