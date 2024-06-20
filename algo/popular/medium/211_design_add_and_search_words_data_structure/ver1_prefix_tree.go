package main

import "fmt"

// TrieNode представляет узел в префиксном дереве
type TrieNode struct {
	children map[rune]*TrieNode // Карта для хранения дочерних узлов
	isEnd    bool               // Флаг, указывающий на конец слова
}

// WordDictionary представляет структуру данных для добавления и поиска слов
type WordDictionary struct {
	root *TrieNode // Корневой узел префиксного дерева
}

// Constructor инициализирует структуру данных WordDictionary
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// AddWord добавляет слово в структуру данных
// Временная сложность: O(m), где m — длина слова
// Пространственная сложность: O(m)
func (wd *WordDictionary) AddWord(word string) {
	node := wd.root // Начинаем с корневого узла
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			// Если символа нет в дочерних узлах, создаем новый узел
			node.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[char] // Переходим к следующему узлу
	}
	node.isEnd = true // Помечаем последний узел как конец слова
}

// Search ищет слово в структуру данных, учитывая возможные точки в шаблоне
// Временная сложность: O(n * 26^m) в худшем случае, где n — количество слов, m — длина слова
// Пространственная сложность: O(m)
func (wd *WordDictionary) Search(word string) bool {
	return searchInNode(wd.root, word, 0)
}

// searchInNode рекурсивно ищет слово в узле префиксного дерева
// Временная сложность: O(26^m) в худшем случае, где m — оставшаяся длина слова
// Пространственная сложность: O(m)
func searchInNode(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		// Если достигли конца слова, проверяем, является ли текущий узел концом слова
		return node.isEnd
	}
	char := rune(word[index])
	if char == '.' {
		// Если текущий символ - точка, проверяем все дочерние узлы
		for _, child := range node.children {
			if searchInNode(child, word, index+1) {
				return true
			}
		}
		return false
	}

	if child, exists := node.children[char]; exists {
		// Если символ существует, продолжаем поиск в дочернем узле
		return searchInNode(child, word, index+1)
	}

	return false
}

func main() {
	wordDict := Constructor()
	wordDict.AddWord("bad")
	wordDict.AddWord("dad")
	wordDict.AddWord("mad")
	fmt.Println(wordDict.Search("pad")) // false
	fmt.Println(wordDict.Search("bad")) // true
	fmt.Println(wordDict.Search(".ad")) // true
	fmt.Println(wordDict.Search("b..")) // true
}
