package main

import "fmt"

// Node - узел бинарного дерева поиска
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// Insert - вставка нового узла в дерево
// TIME COMPLEXITY: O(h), где h - высота дерева. В худшем случае, когда дерево является несбалансированным, h может быть равно n, г
// де n - количество узлов в дереве. В среднем, если дерево сбалансировано, h будет логарифмическим от n.
// SPACE COMPLEXITY: O(h), где h - высота дерева. В худшем случае, когда дерево является несбалансированным,
// h может быть равно n, где n - количество узлов в дереве. В среднем, если дерево сбалансировано, h будет логарифмическим от n.
func (n *Node) Insert(value int) {
	if value < n.Value {
		// Вставка в левое поддерево
		if n.Left == nil {
			n.Left = &Node{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else if value > n.Value {
		// Вставка в правое поддерево
		if n.Right == nil {
			n.Right = &Node{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

// Search - поиск узла в дереве.
// TIME COMPLEXITY: O(h), где h - высота дерева. В худшем случае, когда дерево является несбалансированным, h может быть равно n, г
// где n - количество узлов в дереве. В среднем, если дерево сбалансировано, h будет логарифмическим от n.
// SPACE COMPLEXITY: O(h), где h - высота дерева. В худшем случае, когда дерево является несбалансированным,
// h может быть равно n, где n - количество узлов в дереве. В среднем, если дерево сбалансировано, h будет логарифмическим от n.
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

// PrintInOrder - вывод дерева в отсортированном порядке.
// TIME COMPLEXITY: O(n), где n - количество узлов в дереве. Это происходит, потому что мы посещаем каждый узел только один раз.
// Space ocmplexity: O(h), где h - высота дерева. В худшем случае, когда дерево является несбалансированным,
// h может быть равно n, где n - количество узлов в дереве. В среднем, если дерево сбалансировано, h будет логарифмическим от n.
func (n *Node) PrintInOrder() {
	if n == nil {
		return
	}

	n.Left.PrintInOrder()
	fmt.Println(n.Value)
	n.Right.PrintInOrder()
}

func main() {
	// Создание нового бинарного дерева поиска
	tree := &Node{Value: 100}

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
