package main

import "testing"

func TestTrieSearchWithWildcard(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		queries  []string
		expected []bool
	}{
		{
			name:     "three letter words",
			words:    []string{"bad", "dad", "mad"},
			queries:  []string{"bad", ".ad", "b..", "b.d", "...", "....", "pad", "."},
			expected: []bool{true, true, true, true, true, false, false, false},
		},
		{
			name:     "single word",
			words:    []string{"go"},
			queries:  []string{"go", "g.", ".o", "..", "..."},
			expected: []bool{true, true, true, true, false},
		},
		{
			name:     "empty word",
			words:    []string{""},
			queries:  []string{"", ".", "a"},
			expected: []bool{true, false, false},
		},
		{
			name:     "empty trie",
			words:    []string{},
			queries:  []string{"", ".", "a"},
			expected: []bool{false, false, false},
		},
		{
			name:     "unicode letters",
			words:    []string{"кот", "кит"},
			queries:  []string{"кот", "к.т", ".от", "..т", "...."},
			expected: []bool{true, true, true, true, false},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			trie := NewTrie()
			for _, word := range testCase.words {
				trie.Insert(word)
			}

			for index, query := range testCase.queries {
				if actual := trie.SearchWithWildcard(query); actual != testCase.expected[index] {
					t.Errorf("SearchWithWildcard(%q) = %v, want %v", query, actual, testCase.expected[index])
				}
			}
		})
	}
}
