package main

import (
	"testing"
)

func TestMaxRepeatingGreedy(t *testing.T) {
	tests := []struct {
		name     string
		sequence string
		word     string
		expected int
	}{
		{
			name:     "Example 1: ababc, ab → 2",
			sequence: "ababc",
			word:     "ab",
			expected: 2,
		},
		{
			name:     "Example 2: ababc, ba → 1",
			sequence: "ababc",
			word:     "ba",
			expected: 1,
		},
		{
			name:     "Example 3: ababc, ac → 0",
			sequence: "ababc",
			word:     "ac",
			expected: 0,
		},
		{
			name:     "Empty word",
			sequence: "abc",
			word:     "",
			expected: 0,
		},
		{
			name:     "Single char match",
			sequence: "aaaa",
			word:     "a",
			expected: 4,
		},
		{
			name:     "No match",
			sequence: "hello",
			word:     "world",
			expected: 0,
		},
		{
			name:     "Full repeat",
			sequence: "abcabcabc",
			word:     "abc",
			expected: 3,
		},
		{
			name:     "Partial repeat",
			sequence: "abcabcab",
			word:     "abc",
			expected: 2,
		},
		{
			name:     "Word longer than sequence",
			sequence: "ab",
			word:     "abc",
			expected: 0,
		},
		{
			name:     "Equal length, exact match",
			sequence: "test",
			word:     "test",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxRepeating(tt.sequence, tt.word)
			if result != tt.expected {
				t.Errorf("maxRepeating(%q, %q) = %d, want %d",
					tt.sequence, tt.word, result, tt.expected)
			}
		})
	}
}
