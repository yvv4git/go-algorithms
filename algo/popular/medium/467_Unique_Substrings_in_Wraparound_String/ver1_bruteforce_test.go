package main

import "testing"

func TestFindSubstringInWraproundString(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1: single character",
			s:        "a",
			expected: 1,
		},
		{
			name:     "Example 2: cac",
			s:        "cac",
			expected: 2,
		},
		{
			name:     "Example 3: zab",
			s:        "zab",
			expected: 6,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Sequential letters",
			s:        "abc",
			expected: 6, // "a", "b", "c", "ab", "bc", "abc"
		},
		{
			name:     "Wrap around",
			s:        "zab",
			expected: 6, // "z", "a", "b", "za", "ab", "zab"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSubstringInWraproundString(tt.s)
			if result != tt.expected {
				t.Errorf("findSubstringInWraproundString(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
