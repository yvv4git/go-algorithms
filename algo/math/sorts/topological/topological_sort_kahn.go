//go:build ignore

package main

import (
	"fmt"
)

// Граф представлен в виде списка смежности
type Graph struct {
	vertices int
	adjList  map[int][]int
}

// Создание нового графа
// Time Complexity: O(1)
// Space Complexity: O(1)
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

// Добавление ребра в граф
// Time Complexity: O(1)
// Space Complexity: O(1)
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
}

// Топологическая сортировка с использованием алгоритма Кана
// Time Complexity: O(V + E)
// Space Complexity: O(V)
func topologicalSortKahn(g *Graph) []int {
	// Входящая степень для каждой вершины
	inDegree := make([]int, g.vertices)

	// Вычисляем входящую степень для каждой вершины
	for _, neighbors := range g.adjList {
		for _, neighbor := range neighbors {
			inDegree[neighbor]++
		}
	}

	// Очередь для хранения вершин с входящей степенью 0
	queue := []int{}

	// Добавляем все вершины с входящей степенью 0 в очередь
	for i := 0; i < g.vertices; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Результат топологической сортировки
	result := []int{}

	// Обрабатываем вершины из очереди
	for len(queue) > 0 {
		// Извлекаем вершину из очереди
		u := queue[0]
		queue = queue[1:]

		// Добавляем вершину в результат
		result = append(result, u)

		// Уменьшаем входящую степень соседей вершины u
		for _, neighbor := range g.adjList[u] {
			inDegree[neighbor]--

			// Если входящая степень соседа стала 0, добавляем его в очередь
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Если количество вершин в результате не равно общему количеству вершин,
	// то граф содержит цикл (не является DAG)
	if len(result) != g.vertices {
		fmt.Println("Граф содержит цикл, топологическая сортировка невозможна")
		return nil
	}

	return result
}

func main() {
	// Пример графа
	g := NewGraph(6)
	g.AddEdge(5, 2)
	g.AddEdge(5, 0)
	g.AddEdge(4, 0)
	g.AddEdge(4, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)

	// Топологическая сортировка с использованием алгоритма Кана
	result := topologicalSortKahn(g)
	fmt.Println("Топологическая сортировка (алгоритм Кана):", result)
}
