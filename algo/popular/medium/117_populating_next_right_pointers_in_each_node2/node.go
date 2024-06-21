package main

// Node - определение структуры для узла бинарного дерева
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
