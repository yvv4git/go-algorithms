package matchstickstosquare

import (
	"testing"
)

func TestMakesquareGreedy(t *testing.T) {
	tests := []struct {
		name        string
		matchsticks []int
		expected    bool
	}{
		{
			name:        "Example 1",
			matchsticks: []int{1, 1, 2, 2, 2},
			expected:    true,
		},
		{
			name:        "Example 2",
			matchsticks: []int{3, 3, 3, 3, 4},
			expected:    false,
		},
		{
			name:        "Empty array",
			matchsticks: []int{},
			expected:    false,
		},
		{
			name:        "Single matchstick",
			matchsticks: []int{1},
			expected:    false,
		},
		{
			name:        "Three matchsticks",
			matchsticks: []int{1, 1, 1},
			expected:    false,
		},
		{
			name:        "Four equal matchsticks",
			matchsticks: []int{1, 1, 1, 1},
			expected:    true,
		},
		{
			name:        "Cannot form square - sum not divisible by 4",
			matchsticks: []int{1, 2, 3, 4, 5},
			expected:    false,
		},
		{
			name:        "Matchstick longer than side",
			matchsticks: []int{5, 1, 1, 1, 1, 1, 1},
			expected:    false,
		},
		{
			name:        "All equal - can form square",
			matchsticks: []int{2, 2, 2, 2, 2, 2, 2, 2},
			expected:    true,
		},
		{
			name:        "Large matchsticks",
			matchsticks: []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
			expected:    true,
		},
		{
			name:        "Complex case - sum=120, side=30",
			matchsticks: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			expected:    true, // sum = 120, 120/4 = 30
		},
		{
			name:        "Complex case 2 - sum=20, side=5",
			matchsticks: []int{1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
			expected:    true, // sum = 20, 20/4 = 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := makesquareGreedy(tt.matchsticks)
			if result != tt.expected {
				t.Errorf("makesquareDFS(%v) = %v, want %v", tt.matchsticks, result, tt.expected)
			}
		})
	}
}
