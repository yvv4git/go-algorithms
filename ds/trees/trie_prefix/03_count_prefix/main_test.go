package main

import "testing"

func TestTrieCountWordsWithPrefix(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		queries  []string
		expected []int
	}{
		{
			name:     "dictionary words",
			words:    []string{"app", "apple", "apply", "ape", "bat"},
			queries:  []string{"app", "appl", "ap", "b", "", "z"},
			expected: []int{3, 2, 4, 1, 5, 0},
		},
		{
			name:     "duplicate words",
			words:    []string{"app", "app", "apple"},
			queries:  []string{"app", "appl", ""},
			expected: []int{2, 1, 2},
		},
		{
			name:     "empty trie",
			words:    []string{},
			queries:  []string{"", "a"},
			expected: []int{0, 0},
		},
		{
			name:     "unicode letters",
			words:    []string{"привет", "пример", "приз"},
			queries:  []string{"при", "прив", "про", ""},
			expected: []int{3, 1, 0, 3},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range testCase.words {
				trie.Insert(word)
			}

			for index, query := range testCase.queries {
				if actual := trie.CountWordsWithPrefix(query); actual != testCase.expected[index] {
					t.Errorf("CountWordsWithPrefix(%q) = %d, want %d", query, actual, testCase.expected[index])
				}
			}
		})
	}
}
