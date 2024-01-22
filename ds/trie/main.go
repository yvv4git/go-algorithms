package main

import "fmt"

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return true
}

func main() {
	trie := NewTrie()

	trie.Insert("hello")
	trie.Insert("world")

	fmt.Println(trie.Search("hello"))  // Выведет: true
	fmt.Println(trie.Search("world"))  // Выведет: true
	fmt.Println(trie.Search("hi"))     // Выведет: false
	fmt.Println(trie.StartsWith("he")) // Выведет: true
	fmt.Println(trie.StartsWith("wo")) // Выведет: true
	fmt.Println(trie.StartsWith("hi")) // Выведет: false
}
