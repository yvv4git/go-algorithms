package main

import (
	"reflect"
	"testing"
)

func TestGetLongestSubsequenceDP(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		groups   []int
		expected []string
	}{
		{
			name:     "Example 1: e,a,b with groups 0,0,1",
			words:    []string{"e", "a", "b"},
			groups:   []int{0, 0, 1},
			expected: []string{"e", "b"},
		},
		{
			name:     "Example 2: a,b,c,d with groups 1,0,1,1",
			words:    []string{"a", "b", "c", "d"},
			groups:   []int{1, 0, 1, 1},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "Single element",
			words:    []string{"a"},
			groups:   []int{0},
			expected: []string{"a"},
		},
		{
			name:     "All same groups",
			words:    []string{"a", "b", "c"},
			groups:   []int{0, 0, 0},
			expected: []string{"a"},
		},
		{
			name:     "Alternating groups",
			words:    []string{"a", "b", "c", "d"},
			groups:   []int{0, 1, 0, 1},
			expected: []string{"a", "b", "c", "d"},
		},
		{
			name:     "Two elements same group",
			words:    []string{"a", "b"},
			groups:   []int{1, 1},
			expected: []string{"a"},
		},
		{
			name:     "Two elements different groups",
			words:    []string{"a", "b"},
			groups:   []int{0, 1},
			expected: []string{"a", "b"},
		},
		{
			name:     "Empty input",
			words:    []string{},
			groups:   []int{},
			expected: []string{},
		},
		{
			name:     "Long sequence with pattern",
			words:    []string{"a", "b", "c", "d", "e", "f"},
			groups:   []int{0, 0, 1, 1, 0, 0},
			expected: []string{"a", "c", "e"},
		},
		{
			name:     "All ones then zeros",
			words:    []string{"a", "b", "c", "d", "e"},
			groups:   []int{1, 1, 1, 0, 0},
			expected: []string{"a", "d"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getLongestSubsequenceDP(tt.words, tt.groups)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("getLongestSubsequenceDP(%v, %v) = %v, want %v",
					tt.words, tt.groups, result, tt.expected)
			}
		})
	}
}
