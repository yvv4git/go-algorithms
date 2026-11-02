package main

import "fmt"

// TrieNode - узел префиксного дерева.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie - префиксное дерево.
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

// StartsWith проверяет, есть ли в дереве слово с указанным началом.
// Флаг isEnd не проверяется: префикс не обязан быть целым словом.
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root

	for _, char := range prefix {
		next, exists := node.children[char]
		if !exists {
			return false
		}

		node = next
	}

	return true
}

func main() {
	trie := NewTrie()
	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("banana")

	fmt.Println(trie.StartsWith("app"))  // true
	fmt.Println(trie.StartsWith("appl")) // true
	fmt.Println(trie.StartsWith("ban"))  // true
	fmt.Println(trie.StartsWith("band")) // false
	fmt.Println(trie.StartsWith(""))     // true
}
