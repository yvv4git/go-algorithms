package main

import "fmt"

// TrieNode - один узел префиксного дерева.
// children хранит переходы по символам, а isEnd отмечает конец слова.
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

// Trie - префиксное дерево.
// Корень соответствует пустому префиксу.
type Trie struct {
	root *TrieNode
}

// NewTrie создаёт пустое префиксное дерево.
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert добавляет слово в дерево.
// Для каждого символа создаётся узел, если его ещё нет.
// В конце пути ставится флаг isEnd.
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

// Search проверяет, есть ли слово в дереве целиком.
// Просто найти путь по символам недостаточно: узел должен быть концом слова.
func (t *Trie) Search(word string) bool {
	node := t.root

	for _, char := range word {
		next, exists := node.children[char]
		if !exists {
			return false
		}

		node = next
	}

	return node.isEnd
}

func main() {
	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("app")

	fmt.Println(trie.Search("app"))    // true
	fmt.Println(trie.Search("apple"))  // true
	fmt.Println(trie.Search("appl"))   // false
	fmt.Println(trie.Search("apples")) // false
}
