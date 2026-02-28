package main

import "testing"

func TestFindSubstringInWraproundStringDP2(t *testing.T) {
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
			expected: 6,
		},
		{
			name:     "Wrap around",
			s:        "zab",
			expected: 6,
		},
		{
			name:     "All letters",
			s:        "abcdefghijklmnopqrstuvwxyz",
			expected: 351,
		},
		{
			name:     "Repeated characters",
			s:        "aa",
			expected: 1,
		},
		{
			name:     "Repeated with continuation",
			s:        "aab",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSubstringInWraproundStringDP2(tt.s)
			if result != tt.expected {
				t.Errorf("findSubstringInWraproundStringDP2(%q) = %d, want %d", tt.s, result, tt.expected)
			}
		})
	}
}
