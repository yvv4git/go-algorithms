package main

import "fmt"

// TrieNode - узел префиксного дерева.
// prefixCount хранит число слов, которые проходят через этот узел.
type TrieNode struct {
	children    map[rune]*TrieNode
	prefixCount int
	isEnd       bool
}

// Trie - префиксное дерево со счётчиком слов на каждом префиксе.
type Trie struct {
	root *TrieNode
}

// NewTrie создаёт пустое префиксное дерево.
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert добавляет слово в дерево.
// Возвращает false, если слово уже было добавлено раньше.
func (t *Trie) Insert(word string) bool {
	node := t.root
	path := make([]*TrieNode, 0, len(word)+1)
	path = append(path, node)

	for _, char := range word {
		next, exists := node.children[char]
		if !exists {
			next = &TrieNode{children: make(map[rune]*TrieNode)}
			node.children[char] = next
		}

		node = next
		path = append(path, node)
	}

	if node.isEnd {
		return false
	}

	node.isEnd = true
	for _, pathNode := range path {
		pathNode.prefixCount++
	}

	return true
}

// CountWordsWithPrefix возвращает число слов, которые начинаются с prefix.
func (t *Trie) CountWordsWithPrefix(prefix string) int {
	node := t.root

	for _, char := range prefix {
		next, exists := node.children[char]
		if !exists {
			return 0
		}

		node = next
	}

	return node.prefixCount
}

func main() {
	trie := NewTrie()
	trie.Insert("app")
	trie.Insert("apple")
	trie.Insert("apply")
	trie.Insert("ape")
	trie.Insert("bat")
	trie.Insert("app")

	fmt.Println(trie.CountWordsWithPrefix("app"))  // 3
	fmt.Println(trie.CountWordsWithPrefix("appl")) // 2
	fmt.Println(trie.CountWordsWithPrefix("ap"))   // 4
	fmt.Println(trie.CountWordsWithPrefix("b"))    // 1
	fmt.Println(trie.CountWordsWithPrefix(""))     // 5
	fmt.Println(trie.CountWordsWithPrefix("z"))    // 0
}
