package main

import "fmt"

// Node - узел черно-белого дерева
type Node struct {
	Value  int
	Color  bool // true - черный, false - красный
	Left   *Node
	Right  *Node
	Parent *Node
}

// Insert - вставка нового узла в дерево
// Временная сложность: O(h), где h - высота дерева
// Пространственная сложность: O(1)
func (n *Node) Insert(value int) {
	if n == nil {
		n = &Node{Value: value, Color: true}
	} else if value < n.Value {
		// Вставка в левое поддерево
		if n.Left == nil {
			n.Left = &Node{Value: value, Color: false, Parent: n}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		// Вставка в правое поддерево
		if n.Right == nil {
			n.Right = &Node{Value: value, Color: false, Parent: n}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search - поиск узла в дереве
// Временная сложность: O(h), где h - высота дерева
// Пространственная сложность: O(1)
func (n *Node) Search(value int) bool {
	if n == nil {
		return false
	}

	if value < n.Value {
		return n.Left.Search(value)
	} else if value > n.Value {
		return n.Right.Search(value)
	}

	return true
}

// PrintInOrder - вывод дерева в отсортированном порядке
// Временная сложность: O(n), где n - количество узлов в дереве
// Пространственная сложность: O(n)
func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}

	n.Left.PrintInOrder()
	fmt.Println(n.Value)
	n.Right.PrintInOrder()
}

func main() {
	// Создание нового черно-белого дерева
	tree := &Node{Value: 100, Color: true}

	// Вставка узлов в дерево
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(50)
	tree.Insert(150)

	// Вывод дерева в отсортированном порядке
	tree.PrintInOrder()

	// Поиск узла в дереве
	fmt.Println(tree.Search(150)) // true
	fmt.Println(tree.Search(400)) // false
}
