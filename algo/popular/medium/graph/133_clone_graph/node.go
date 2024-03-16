package main

// Node определяет структуру узла графа.
type Node struct {
	Val       int
	Neighbors []*Node
}
