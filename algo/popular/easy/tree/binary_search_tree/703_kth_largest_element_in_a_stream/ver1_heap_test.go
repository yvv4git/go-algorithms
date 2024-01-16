package main

import (
	"fmt"
	"testing"
)

func TestKthLargest(t *testing.T) {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})

	tests := []struct {
		val    int
		expect int
	}{
		{3, 4},
		{5, 5},
		{10, 5},
		{9, 8},
		{4, 8},
	}

	for _, test := range tests {
		result := kthLargest.Add(test.val)
		if result != test.expect {
			t.Errorf("Expected %d, but got %d", test.expect, result)
		} else {
			fmt.Printf("Test passed with input %d\n", test.val)
		}
	}
}
