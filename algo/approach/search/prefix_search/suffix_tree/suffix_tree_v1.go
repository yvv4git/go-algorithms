package main

import "fmt"

type suffixTreeNode struct {
	children map[byte]*suffixTreeNode
	start    int
	end      *int
	suffixID int
}

func newSuffixTreeNode(start int, end *int) *suffixTreeNode {
	return &suffixTreeNode{
		children: make(map[byte]*suffixTreeNode),
		start:    start,
		end:      end,
		suffixID: -1,
	}
}

type suffixTree struct {
	text string
	root *suffixTreeNode
}

func NewSuffixTree(text string) *suffixTree {
	/*
		METHOD: Построение суффиксного дерева (наивное, O(n²))
		TIME COMPLEXITY: O(n²) — наивная вставка каждого суффикса
		SPACE COMPLEXITY: O(n²) в худшем случае

		Суффиксное дерево — сжатый бор всех суффиксов строки.
		Каждый путь от корня к листу соответствует суффиксу.
		Ребра могут содержать подстроки (не только один символ),
		что отличает его от обычного бора.

		НА ЗАМЕТКУ: алгоритм Укконена строит дерево за O(n),
		но для демонстрации сути используется наивный вариант.
	*/
	st := &suffixTree{
		text: text,
		root: newSuffixTreeNode(-1, nil),
	}

	for i := 0; i < len(text); i++ {
		st.insertSuffix(i)
	}

	return st
}

func (st *suffixTree) insertSuffix(start int) {
	node := st.root

	for i := start; i < len(st.text); {
		ch := st.text[i]

		child, exists := node.children[ch]
		if !exists {
			end := len(st.text)
			newNode := newSuffixTreeNode(i, &end)
			newNode.suffixID = start
			node.children[ch] = newNode
			return
		}

		j := child.start
		for j < *child.end && i < len(st.text) && st.text[j] == st.text[i] {
			j++
			i++
		}

		if j == *child.end {
			node = child
		} else {
			splitEnd := j
			split := newSuffixTreeNode(child.start, &splitEnd)
			child.start = j
			split.children[st.text[j]] = child

			end := len(st.text)
			leaf := newSuffixTreeNode(i, &end)
			leaf.suffixID = start
			split.children[st.text[i]] = leaf

			node.children[ch] = split
			return
		}
	}
}

func (st *suffixTree) Search(pattern string) []int {
	/*
		METHOD: Поиск подстроки в суффиксном дереве
		TIME COMPLEXITY: O(m), где m — длина паттерна
		SPACE COMPLEXITY: O(1)

		Спуск по дереву по символам паттерна. Если паттерн прочитан —
		собираем все суффиксные ID в поддереве.
	*/
	if len(pattern) == 0 {
		return nil
	}

	node := st.root
	i := 0

	for i < len(pattern) {
		ch := pattern[i]
		child, exists := node.children[ch]
		if !exists {
			return nil
		}

		j := child.start
		for j < *child.end && i < len(pattern) && st.text[j] == pattern[i] {
			j++
			i++
		}

		if i == len(pattern) {
			var matches []int
			st.collectLeafIDs(child, &matches)
			return matches
		}

		if j < *child.end {
			return nil
		}

		node = child
	}

	return nil
}

func (st *suffixTree) collectLeafIDs(node *suffixTreeNode, ids *[]int) {
	if node.suffixID != -1 {
		*ids = append(*ids, node.suffixID)
	}
	for _, child := range node.children {
		st.collectLeafIDs(child, ids)
	}
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	tree := NewSuffixTree(text)
	positions := tree.Search(pattern)

	fmt.Printf("Текст:    %s\n", text)
	fmt.Printf("Паттерн:  %s\n", pattern)
	if len(positions) > 0 {
		fmt.Printf("Найден на позициях: %v\n", positions)
	} else {
		fmt.Println("Паттерн не найден в тексте.")
	}
}
