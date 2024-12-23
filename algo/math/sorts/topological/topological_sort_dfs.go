//go:build ignore

package main

import (
	"fmt"
)

// Граф представлен в виде списка смежности
type Graph struct {
	// vertices - количество вершин в графе
	vertices int

	// adjList - список смежности, где ключ — это вершина,
	// а значение — список вершин, в которые можно попасть из данной вершины
	adjList map[int][]int
}

// Создание нового графа
// Time Complexity: O(1) - константное время для инициализации структуры
// Space Complexity: O(1) - константное пространство для хранения структуры
func NewGraph(vertices int) *Graph {
	// Инициализация графа с заданным количеством вершин
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

// Добавление ребра в граф
// Time Complexity: O(1) - константное время для добавления элемента в список
// Space Complexity: O(1) - константное дополнительное пространство
func (g *Graph) AddEdge(u, v int) {
	// Добавляем ребро из вершины u в вершину v
	// Это означает, что из u можно попасть в v
	g.adjList[u] = append(g.adjList[u], v)
}

// Топологическая сортировка с использованием DFS
// Time Complexity: O(V + E) - где V — количество вершин, E — количество ребер
// Space Complexity: O(V) - для хранения стека и массива visited
func topologicalSort(g *Graph) []int {
	// visited - массив для отслеживания посещенных вершин
	visited := make([]bool, g.vertices)

	// stack - стек для хранения результата топологической сортировки
	stack := []int{}

	// Вспомогательная функция для DFS
	var dfs func(int)
	dfs = func(v int) {
		// Отмечаем текущую вершину как посещенную
		visited[v] = true

		// Рекурсивно обходим всех соседей текущей вершины
		for _, neighbor := range g.adjList[v] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}

		// После завершения обхода всех соседей добавляем вершину в стек
		stack = append(stack, v)
	}

	// Запускаем DFS для всех вершин
	for i := 0; i < g.vertices; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	// Результат топологической сортировки — стек в обратном порядке
	result := make([]int, 0, len(stack))
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, stack[i])
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

	// Топологическая сортировка
	result := topologicalSort(g)
	fmt.Println("Топологическая сортировка:", result)
}
