package main

import "fmt"

// TrieNode - узел префиксного дерева.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie - префиксное дерево с поддержкой шаблона.
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

// SearchWithWildcard проверяет, есть ли слово, совпадающее с шаблоном.
// Точка в шаблоне совпадает с любой одной буквой.
func (t *Trie) SearchWithWildcard(word string) bool {
	return t.root.matches([]rune(word), 0)
}

// matches рекурсивно сравнивает шаблон с путями в дереве.
// index - позиция текущего символа в шаблоне.
func (node *TrieNode) matches(pattern []rune, index int) bool {
	if index == len(pattern) {
		return node.isEnd
	}

	char := pattern[index]

	if char == '.' {
		for _, child := range node.children {
			if child.matches(pattern, index+1) {
				return true
			}
		}

		return false
	}

	child, exists := node.children[char]
	if !exists {
		return false
	}

	return child.matches(pattern, index+1)
}

func main() {
	trie := NewTrie()
	trie.Insert("bad")
	trie.Insert("dad")
	trie.Insert("mad")

	fmt.Println(trie.SearchWithWildcard("bad"))  // true
	fmt.Println(trie.SearchWithWildcard(".ad"))  // true
	fmt.Println(trie.SearchWithWildcard("b.."))  // true
	fmt.Println(trie.SearchWithWildcard("b.d"))  // true
	fmt.Println(trie.SearchWithWildcard("..."))  // true
	fmt.Println(trie.SearchWithWildcard("....")) // false
	fmt.Println(trie.SearchWithWildcard("pad"))  // false
	fmt.Println(trie.SearchWithWildcard("."))    // false
}
