//go:build ignore

package main

import (
	"fmt"
	"math"
)

// Edge представляет ребро графа
type Edge struct {
	To     string  // Вершина, в которую ведет ребро
	Weight float64 // Вес ребра
}

// Graph представляет граф в виде списка смежности
type Graph map[string][]Edge

// Dijkstra реализует алгоритм Дейкстры с использованием массива
func Dijkstra(graph Graph, start string) map[string]float64 {
	/*
		METHOD: Dijkstra's Algorithm with Array
		1. Инициализируем массив расстояний: dist[start] = 0, остальные = ∞
		2. Инициализируем массив посещенных вершин
		3. В цикле находим непосещенную вершину с минимальным расстоянием
		4. Помечаем ее как посещенную и обновляем расстояния до соседей
		5. Повторяем до обработки всех вершин

		TIME COMPLEXITY: O(V^2) (V - вершины, сканирование массива для поиска минимума)
		SPACE COMPLEXITY: O(V) (массивы расстояний и посещенных вершин)
	*/

	dist := make(map[string]float64) // Словарь расстояний до вершин
	visited := make(map[string]bool) // Словарь посещенных вершин

	// Инициализация: устанавливаем расстояние до начальной вершины в 0,
	// а до всех остальных вершин в бесконечность
	for vertex := range graph {
		if vertex == start {
			dist[vertex] = 0
		} else {
			dist[vertex] = math.Inf(1)
		}
		visited[vertex] = false
	}

	// Основной цикл алгоритма Дейкстры
	for i := 0; i < len(graph); i++ {
		// Находим вершину с минимальным расстоянием среди непосещенных
		var minVertex string
		minDist := math.Inf(1)
		for vertex, d := range dist {
			if !visited[vertex] && d < minDist {
				minDist = d
				minVertex = vertex
			}
		}

		if minVertex == "" {
			break // Все вершины обработаны
		}

		visited[minVertex] = true

		// Обновляем расстояния до соседей
		for _, edge := range graph[minVertex] {
			neighbor := edge.To
			if !visited[neighbor] {
				newDist := dist[minVertex] + edge.Weight
				if newDist < dist[neighbor] {
					dist[neighbor] = newDist
				}
			}
		}
	}

	return dist
}

func main() {
	// Пример графа в виде словаря смежности
	graph := Graph{
		"A": {{"B", 1}, {"C", 4}},
		"B": {{"A", 1}, {"C", 2}, {"D", 5}},
		"C": {{"A", 4}, {"B", 2}, {"D", 1}},
		"D": {{"B", 5}, {"C", 1}},
	}

	startVertex := "A"                                // Начальная вершина
	shortestDistances := Dijkstra(graph, startVertex) // Запуск алгоритма Дейкстры
	fmt.Println(shortestDistances)                    // Вывод результатов map[A:0 B:1 C:3 D:4]
}
