package main

import "testing"

func TestTrieSearch(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		queries  []string
		expected []bool
	}{
		{
			name:     "word and its prefix",
			words:    []string{"apple", "app"},
			queries:  []string{"app", "apple", "appl", ""},
			expected: []bool{true, true, false, false},
		},
		{
			name:     "single word",
			words:    []string{"go"},
			queries:  []string{"go", "g", "golang"},
			expected: []bool{true, false, false},
		},
		{
			name:     "duplicate insert",
			words:    []string{"go", "go"},
			queries:  []string{"go"},
			expected: []bool{true},
		},
		{
			name:     "empty word",
			words:    []string{""},
			queries:  []string{"", "a"},
			expected: []bool{true, false},
		},
		{
			name:     "unicode letters",
			words:    []string{"привет"},
			queries:  []string{"привет", "при", ""},
			expected: []bool{true, false, false},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range testCase.words {
				trie.Insert(word)
			}

			for index, query := range testCase.queries {
				if actual := trie.Search(query); actual != testCase.expected[index] {
					t.Errorf("Search(%q) = %v, want %v", query, actual, testCase.expected[index])
				}
			}
		})
	}
}
