package main

import (
	"container/list"
	"fmt"
)

// GraphV2 represents an adjacency list graph.
type GraphV2 struct {
	Vertices int
	adjList  map[int][]int
}

// addEdge adds an edge to the graph.
func (g *GraphV2) addEdge(v int, w int) {
	g.adjList[v] = append(g.adjList[v], w)
}

// BFS выполняет поиск в ширину в графе.
func (g *GraphV2) BFS(start int) {
	/*
		METHOD: Stack
		TIME COMPLEXITY: O(V + E), где V - количество вершин, E - количество ребер. Это связано с тем, что в худшем случае мы посетим каждую вершину и каждое ребро графа.
		SPACE COMPLEXITY: O(V), так как в худшем случае мы храним в очереди все вершины графа. Это происходит, когда граф является несвязным и мы должны посетить все вершины, прежде чем достигнуть связности.
	*/
	// Создаем очередь для BFS и помещаем в нее стартовую вершину.
	queue := list.New()
	queue.PushBack(start)

	// Помечаем стартовую вершину как посещенную.
	visited := make(map[int]bool)
	visited[start] = true

	// Пока очередь не пуста, выполняем следующие действия.
	for queue.Len() > 0 {
		// Извлекаем вершину из очереди.
		front := queue.Front()
		v := front.Value.(int)
		queue.Remove(front)

		// Выводим вершину.
		fmt.Println(v)

		// Получаем все соседние вершины для извлеченной вершины.
		// Если соседняя вершина не была посещена, помечаем ее как посещенную и помещаем в очередь.
		for _, adj := range g.adjList[v] {
			if !visited[adj] {
				visited[adj] = true
				queue.PushBack(adj)
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
//	fmt.Println("Breadth-First Search (starting from vertex 2):")
//	g.BFS(2)
//}
