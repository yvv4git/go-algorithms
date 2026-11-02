package main

import (
	"slices"
	"testing"
)

func TestTrieAutocomplete(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		prefix   string
		expected []string
	}{
		{
			name:     "several matches",
			words:    []string{"app", "apple", "apply", "ape", "bat"},
			prefix:   "app",
			expected: []string{"app", "apple", "apply"},
		},
		{
			name:     "one match",
			words:    []string{"app", "apple", "apply", "ape", "bat"},
			prefix:   "b",
			expected: []string{"bat"},
		},
		{
			name:     "sorted result",
			words:    []string{"app", "apple", "apply", "ape", "bat"},
			prefix:   "ap",
			expected: []string{"ape", "app", "apple", "apply"},
		},
		{
			name:     "all words",
			words:    []string{"app", "apple", "apply", "ape", "bat"},
			prefix:   "",
			expected: []string{"ape", "app", "apple", "apply", "bat"},
		},
		{
			name:     "no matches",
			words:    []string{"app", "apple"},
			prefix:   "z",
			expected: nil,
		},
		{
			name:     "empty trie",
			words:    []string{},
			prefix:   "a",
			expected: nil,
		},
		{
			name:     "unicode letters",
			words:    []string{"привет", "пример", "приз"},
			prefix:   "при",
			expected: []string{"привет", "приз", "пример"},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range testCase.words {
				trie.Insert(word)
			}

			actual := trie.Autocomplete(testCase.prefix)
			if !slices.Equal(actual, testCase.expected) {
				t.Errorf("Autocomplete(%q) = %v, want %v", testCase.prefix, actual, testCase.expected)
			}
		})
	}
}
