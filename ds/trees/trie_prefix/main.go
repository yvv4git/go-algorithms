package main

import "fmt"

// TrieNode - структура для узла дерева.
type TrieNode struct {
	children map[rune]*TrieNode // Дочерние узлы (карта символов -> узел)
	isEnd    bool               // Флаг, указывающий, что узел завершает слово
}

// Trie - структура префиксного дерева.
type Trie struct {
	root *TrieNode // Корень дерева
}

// NewTrie - конструктор для создания нового префиксного дерева.
// Время работы: O(1), так как создается только один корень.
// Пространственная сложность: O(1), так как используется только память под корень.
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert - метод для вставки слова в префиксное дерево.
// Время работы: O(k), где k - длина слова, так как нужно пройти по каждому символу слова.
// Пространственная сложность: O(k), так как в худшем случае для каждого символа будет создан новый узел.
func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		// Если дочернего узла для символа нет, создаем его.
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		// Переходим к следующему узлу.
		node = node.children[char]
	}
	// Отмечаем конец слова.
	node.isEnd = true
}

// Search - метод для поиска точного совпадения слова в префиксном дереве.
// Время работы: O(k), где k - длина слова, так как нужно пройти по каждому символу слова.
// Пространственная сложность: O(1), так как используем только указатели для навигации по дереву.
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		// Если дочернего узла для символа нет, возвращаем false.
		if _, exists := node.children[char]; !exists {
			return false
		}
		// Переходим к следующему узлу.
		node = node.children[char]
	}
	// Если достигли конца слова и узел отмечен как конец, то слово найдено.
	return node.isEnd
}

// StartsWith - метод для проверки, существует ли слово с данным префиксом в дереве.
// Время работы: O(k), где k - длина префикса, так как нужно пройти по каждому символу префикса.
// Пространственная сложность: O(1), так как используем только указатели для навигации по дереву.
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		// Если дочернего узла для символа нет, возвращаем false.
		if _, exists := node.children[char]; !exists {
			return false
		}
		// Переходим к следующему узлу.
		node = node.children[char]
	}
	// Префикс найден, возвращаем true.
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
