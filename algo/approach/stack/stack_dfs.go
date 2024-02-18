package main

import (
	"fmt"
)

// Graph represents an adjacency list graph.
type Graph struct {
	Vertices int
	adjList  map[int][]int
}

// addEdge adds an edge to the graph.
func (g *Graph) addEdge(v int, w int) {
	g.adjList[v] = append(g.adjList[v], w)
}

// DFS performs a depth-first search on the graph.
func (g *Graph) DFS(start int) {
	/*
		METHOD: Stack
		TIME COMPLEXITY: O(V + E) обусловлена тем, что в худшем случае мы посетим каждую вершину и каждое ребро графа.
		SPACE COMPLEXITY: O(V + E) обусловлена тем, что в худшем случае мы посетим каждую вершину и каждое ребро графа.

		DFS (Depth-First Search) - это алгоритм обхода или поиска графа, который использует стек.
	*/
	visited := make(map[int]bool)

	// Создаем стек для DFS и добавляем стартовую вершину.
	var stack []int
	stack = append(stack, start)

	for len(stack) > 0 {
		// Извлекаем вершину из стека.
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Если вершина не посещена, выводим ее и помечаем как посещенную.
		if !visited[v] {
			fmt.Println(v)
			visited[v] = true
		}

		// Получаем все соседние вершины для извлеченной вершины.
		// Если соседняя вершина не была посещена, добавляем ее в стек.
		for _, adj := range g.adjList[v] {
			if !visited[adj] {
				stack = append(stack, adj)
			}
		}
	}
}

//func main() {
//	g := Graph{
//		Vertices: 4,
//		adjList:  make(map[int][]int),
//	}
//
//	g.addEdge(0, 1)
//	g.addEdge(0, 2)
//	g.addEdge(1, 2)
//	g.addEdge(2, 0)
//	g.addEdge(2, 3)
//	g.addEdge(3, 3)
//
//	fmt.Println("Depth-First Search (starting from vertex 2):")
//	g.DFS(2)
//}
