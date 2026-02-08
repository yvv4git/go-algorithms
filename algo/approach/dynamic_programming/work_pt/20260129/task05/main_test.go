package main

import (
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		N, M     int
		expected int
	}{
		{"N=1, M=1", 1, 1, 1},
		{"N=3, M=2", 3, 2, 1},
		{"N=31, M=34", 31, 34, 293930},
		{"N=2, M=2", 2, 2, 0}, // no moves possible
		{"N=1, M=2", 1, 2, 0}, // can't reach
		{"N=2, M=1", 2, 1, 0}, // can't reach
		{"N=4, M=3", 4, 3, 0}, // no way to reach (4,3) with the given moves
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := solve(tt.N, tt.M)
			if result != tt.expected {
				t.Errorf("solve(%d, %d) = %d, want %d", tt.N, tt.M, result, tt.expected)
			}
		})
	}
}
