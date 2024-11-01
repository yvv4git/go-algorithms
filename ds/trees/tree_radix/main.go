package main

import (
	"fmt"
)

// Node представляет узел в radix дереве
type Node struct {
	children map[rune]*Node
	isEnd    bool
}

// RadixTree представляет само radix дерево
type RadixTree struct {
	root *Node
}

// NewRadixTree создает новое radix дерево
func NewRadixTree() *RadixTree {
	return &RadixTree{
		root: &Node{
			children: make(map[rune]*Node),
			isEnd:    false,
		},
	}
}

// Insert вставляет ключ в radix дерево
func (rt *RadixTree) Insert(key string) {
	node := rt.root
	for _, char := range key {
		if _, exists := node.children[char]; !exists {
			node.children[char] = &Node{
				children: make(map[rune]*Node),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

// Search ищет ключ в radix дереве
func (rt *RadixTree) Search(key string) bool {
	node := rt.root
	for _, char := range key {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

// PrintAllKeys выводит все ключи в radix дереве
func (rt *RadixTree) PrintAllKeys() {
	rt.printKeysFromNode(rt.root, []rune{})
}

// printKeysFromNode рекурсивно выводит ключи из узла
func (rt *RadixTree) printKeysFromNode(node *Node, prefix []rune) {
	if node.isEnd {
		fmt.Println(string(prefix))
	}
	for char, child := range node.children {
		rt.printKeysFromNode(child, append(prefix, char))
	}
}

func main() {
	rt := NewRadixTree()

	// Вставляем ключи
	rt.Insert("car")
	rt.Insert("cat")
	rt.Insert("cart")
	rt.Insert("card")

	// Ищем ключи
	fmt.Println("Search 'car':", rt.Search("car"))   // true
	fmt.Println("Search 'cat':", rt.Search("cat"))   // true
	fmt.Println("Search 'cart':", rt.Search("cart")) // true
	fmt.Println("Search 'card':", rt.Search("card")) // true
	fmt.Println("Search 'ca':", rt.Search("ca"))     // false

	// Выводим все ключи
	fmt.Println("All keys in the tree:")
	rt.PrintAllKeys()
}
