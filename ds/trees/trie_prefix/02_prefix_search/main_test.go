package main

import "testing"

func TestTrieStartsWith(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		queries  []string
		expected []bool
	}{
		{
			name:     "dictionary words",
			words:    []string{"app", "apple", "banana"},
			queries:  []string{"app", "appl", "ban", "band", ""},
			expected: []bool{true, true, true, false, true},
		},
		{
			name:     "single word",
			words:    []string{"go"},
			queries:  []string{"g", "go", "gol", "x"},
			expected: []bool{true, true, false, false},
		},
		{
			name:     "empty trie",
			words:    []string{},
			queries:  []string{"", "a"},
			expected: []bool{true, false},
		},
		{
			name:     "unicode letters",
			words:    []string{"привет", "пример"},
			queries:  []string{"при", "прив", "про", ""},
			expected: []bool{true, true, false, true},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range testCase.words {
				trie.Insert(word)
			}

			for index, query := range testCase.queries {
				if actual := trie.StartsWith(query); actual != testCase.expected[index] {
					t.Errorf("StartsWith(%q) = %v, want %v", query, actual, testCase.expected[index])
				}
			}
		})
	}
}
