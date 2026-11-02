package main

import "testing"

func TestTrieLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected string
	}{
		{
			name:     "common prefix",
			words:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "no common prefix",
			words:    []string{"dog", "cat"},
			expected: "",
		},
		{
			name:     "single word",
			words:    []string{"apple"},
			expected: "apple",
		},
		{
			name:     "word is prefix of another",
			words:    []string{"app", "apple"},
			expected: "app",
		},
		{
			name:     "empty set",
			words:    []string{},
			expected: "",
		},
		{
			name:     "empty word in set",
			words:    []string{"apple", ""},
			expected: "",
		},
		{
			name:     "duplicate words",
			words:    []string{"abc", "abc"},
			expected: "abc",
		},
		{
			name:     "unicode words",
			words:    []string{"привет", "пример"},
			expected: "при",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range testCase.words {
				trie.Insert(word)
			}

			if actual := trie.LongestCommonPrefix(); actual != testCase.expected {
				t.Errorf("LongestCommonPrefix() = %q, want %q", actual, testCase.expected)
			}
		})
	}
}
