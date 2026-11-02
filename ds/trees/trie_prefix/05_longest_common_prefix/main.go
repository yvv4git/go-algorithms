package main

import (
	"fmt"
	"strings"
)

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

// LongestCommonPrefix возвращает самый длинный общий префикс всех слов.
// Идём по дереву, пока из узла выходит ровно одно ребро и узел не является концом слова.
func (t *Trie) LongestCommonPrefix() string {
	node := t.root
	var prefix strings.Builder

	for {
		if node.isEnd || len(node.children) != 1 {
			break
		}

		var nextChar rune
		var nextNode *TrieNode

		for char, child := range node.children {
			nextChar = char
			nextNode = child
		}

		prefix.WriteRune(nextChar)
		node = nextNode
	}

	return prefix.String()
}

func main() {
	first := NewTrie()
	first.Insert("flower")
	first.Insert("flow")
	first.Insert("flight")
	fmt.Println(first.LongestCommonPrefix()) // fl

	second := NewTrie()
	second.Insert("dog")
	second.Insert("cat")
	fmt.Println(second.LongestCommonPrefix()) // пустая строка

	third := NewTrie()
	third.Insert("app")
	third.Insert("apple")
	fmt.Println(third.LongestCommonPrefix()) // app
}
