package main

import "fmt"

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{children: make(map[byte]*TrieNode)}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

func (t *Trie) Insert(word string) {
	/*
		METHOD: Вставка в бор
		TIME COMPLEXITY: O(m), где m — длина слова
		SPACE COMPLEXITY: O(m) в худшем случае (если ни один префикс не совпадает)
	*/
	node := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if node.children[ch] == nil {
			node.children[ch] = newTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	/*
		METHOD: Поиск слова в боре
		TIME COMPLEXITY: O(m), где m — длина слова
		SPACE COMPLEXITY: O(1)
	*/
	node := t.root
	for i := 0; i < len(word); i++ {
		ch := word[i]
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	/*
		METHOD: Проверка наличия префикса в боре
		TIME COMPLEXITY: O(m), где m — длина префикса
		SPACE COMPLEXITY: O(1)
	*/
	node := t.root
	for i := 0; i < len(prefix); i++ {
		ch := prefix[i]
		if node.children[ch] == nil {
			return false
		}
		node = node.children[ch]
	}
	return true
}

func main() {
	/*
		Пример использования бора (trie) для хранения и поиска строк.
		Бор — это дерево, где каждый путь от корня к листу представляет
		собой строку (или префикс). Позволяет эффективно искать слова и префиксы.
	*/
	trie := NewTrie()
	words := []string{"hello", "world", "hell", "high", "hill", "worlds"}

	for _, w := range words {
		trie.Insert(w)
	}

	fmt.Println("Слова в боре:", words)
	fmt.Println()

	tests := []struct {
		query string
		op    string
	}{
		{"hello", "search"},
		{"hell", "search"},
		{"heaven", "search"},
		{"hell", "startsWith"},
		{"wor", "startsWith"},
		{"xyz", "startsWith"},
	}

	for _, tt := range tests {
		var ok bool
		if tt.op == "search" {
			ok = trie.Search(tt.query)
		} else {
			ok = trie.StartsWith(tt.query)
		}
		opLabel := map[string]string{"search": "Поиск", "startsWith": "Префикс"}[tt.op]
		if ok {
			fmt.Printf("✓ %s: «%s» найден\n", opLabel, tt.query)
		} else {
			fmt.Printf("✗ %s: «%s» не найден\n", opLabel, tt.query)
		}
	}
}
