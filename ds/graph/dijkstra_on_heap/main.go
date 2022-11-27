package main

import (
	"fmt"
)

func main() {
	fmt.Println("Dijkstra")

	// Example
	graph := NewGraph()
	graph.AddEdge("S", "B", 4)
	graph.AddEdge("S", "C", 2)
	graph.AddEdge("B", "C", 1)
	graph.AddEdge("B", "D", 5)
	graph.AddEdge("C", "D", 8)
	graph.AddEdge("C", "E", 10)
	graph.AddEdge("D", "E", 2)
	graph.AddEdge("D", "T", 6)
	graph.AddEdge("E", "T", 2)

	fmt.Println(graph.GetPath("S", "T"))
}
