package main

import (
	"testing"
)

func TestCountSequences(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{1, 2}, // "0", "1"
		{2, 4}, // "00", "01", "10", "11"
		{3, 7}, // все кроме "111"
		{4, 13},
		{5, 24},
		{35, 2082876103},
	}

	for _, tt := range tests {
		result := countSequences(tt.n)
		if result != tt.expected {
			t.Errorf("countSequences(%d) = %d, want %d", tt.n, result, tt.expected)
		}
	}
}
