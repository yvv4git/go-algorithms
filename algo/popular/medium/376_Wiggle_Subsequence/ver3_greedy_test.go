package wigglesubsequence

import "testing"

func TestWiggleMaxLength_Ver3(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"single element", []int{1}, 1},
		{"two equal elements", []int{1, 1}, 1},
		{"two different elements", []int{1, 7}, 2},
		{"example 1", []int{1, 7, 4, 9, 2, 5}, 6},
		{"example 2", []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{"example 3", []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{"decreasing", []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 2},
		{"alternating", []int{1, 3, 2, 4, 3, 5}, 6},
		{"with equal adjacent", []int{1, 1, 1, 1, 1}, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := wiggleMaxLengthGreedy(tc.input)
			if result != tc.expected {
				t.Errorf("wiggleMaxLengthGreedy(%v) = %d; want %d", tc.input, result, tc.expected)
			}
		})
	}
}