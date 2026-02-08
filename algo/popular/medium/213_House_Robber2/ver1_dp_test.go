package main

import (
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: [2,3,2]",
			nums:     []int{2, 3, 2},
			expected: 3,
		},
		{
			name:     "Example 2: [1,2,3,1]",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "Example 3: [1,2,3]",
			nums:     []int{1, 2, 3},
			expected: 3,
		},
		{
			name:     "Single house",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Two houses",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "Large values",
			nums:     []int{100, 1, 1, 100, 1, 1, 100},
			expected: 201,
		},
		{
			name:     "Circular with small values",
			nums:     []int{1, 3, 1, 2},
			expected: 5,
		},
		{
			name:     "Circular with equal houses",
			nums:     []int{1, 1, 1, 1},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rob(tt.nums)
			if result != tt.expected {
				t.Errorf("rob(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestRobEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Two houses different",
			nums:     []int{5, 10},
			expected: 10,
		},
		{
			name:     "Three houses optimal middle",
			nums:     []int{10, 1, 10},
			expected: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rob(tt.nums)
			if result != tt.expected {
				t.Errorf("rob(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}
