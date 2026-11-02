package main

import (
	"fmt"
	"sort"
)

// TrieNode - узел префиксного дерева.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie - префиксное дерево для автодополнения.
type Trie struct {
	root *TrieNode
}

// NewTrie создаёт пустое префиксное дерево.
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert добавляет слово в дерево.
func (t *Trie) Insert(word string) {
	node := t.root

	for _, char := range word {
		next, exists := node.children[char]
		if !exists {
			next = &TrieNode{children: make(map[rune]*TrieNode)}
			node.children[char] = next
		}

		node = next
	}

	node.isEnd = true
}

// Autocomplete возвращает все слова с указанным префиксом в алфавитном порядке.
func (t *Trie) Autocomplete(prefix string) []string {
	node := t.root

	for _, char := range prefix {
		next, exists := node.children[char]
		if !exists {
			return nil
		}

		node = next
	}

	words := make([]string, 0)
	node.collect(prefix, &words)
	sort.Strings(words)

	return words
}

// collect обходит поддерево и собирает все слова в words.
// Параметр current хранит строку, накопленную от корня до текущего узла.
func (node *TrieNode) collect(current string, words *[]string) {
	if node.isEnd {
		*words = append(*words, current)
	}

	for char, child := range node.children {
		child.collect(current+string(char), words)
	}
}

func main() {
	trie := NewTrie()
	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("apply")
	trie.Insert("ape")
	trie.Insert("bat")

	fmt.Println(trie.Autocomplete("app"))  // [app apple apply]
	fmt.Println(trie.Autocomplete("appl")) // [apple apply]
	fmt.Println(trie.Autocomplete("ap"))   // [ape app apple apply]
	fmt.Println(trie.Autocomplete("b"))    // [bat]
	fmt.Println(trie.Autocomplete("z"))    // []
	fmt.Println(trie.Autocomplete(""))     // [ape app apple apply bat]
}
