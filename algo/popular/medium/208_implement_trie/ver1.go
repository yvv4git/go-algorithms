package main

import "fmt"

// Node - структура для узла Trie
type Node struct {
	children [26]*Node // Для каждой буквы алфавита может быть дочерний узел
	isEnd    bool      // Флаг, указывающий на конец слова
}

// Trie - структура для самого дерева
type Trie struct {
	root *Node // Корень дерева
}

// Конструктор для создания нового узла Trie
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func newNode() *Node {
	node := &Node{}
	node.isEnd = false
	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}
	return node
}

// Constructor Конструктор для создания нового Trie
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func Constructor() Trie {
	trie := Trie{}
	trie.root = newNode()
	return trie
}

// Insert Вставка слова в Trie
// Временная сложность: O(L), где L - длина слова
// Пространственная сложность: O(L)
func (this *Trie) Insert(word string) {
	node := this.root
	for _, ch := range word {
		index := ch - 'a' // Получаем индекс буквы в алфавите
		if node.children[index] == nil {
			node.children[index] = newNode()
		}
		node = node.children[index]
	}
	node.isEnd = true // Помечаем конец слова
}

// Search Поиск слова в Trie
// Временная сложность: O(L), где L - длина слова
// Пространственная сложность: O(1)
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd // Возвращаем true, если слово полностью найдено и оно является конечным узлом
}

// StartsWith Проверка наличия префикса в Trie
// Временная сложность: O(L), где L - длина префикса
// Пространственная сложность: O(1)
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true // Возвращаем true, если префикс найден
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // Выведет: true
	fmt.Println(trie.Search("app"))     // Выведет: false
	fmt.Println(trie.StartsWith("app")) // Выведет: true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // Выведет: true
}
