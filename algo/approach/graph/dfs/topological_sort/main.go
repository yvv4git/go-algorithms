package main

import "fmt"

// Функция для проверки наличия цикла в подграфах
func isCyclicUtil(graph map[string][]string, v string, visited map[string]bool, recStack map[string]bool) bool {
	if recStack[v] {
		return true
	}
	if visited[v] {
		return false
	}

	// Помечаем текущую вершину как посещенную и добавляем ее в стек вызовов
	visited[v] = true
	recStack[v] = true

	// Рекурсивно вызываем для всех соседних вершин
	for _, neighbour := range graph[v] {
		if isCyclicUtil(graph, neighbour, visited, recStack) {
			return true
		}
	}

	// Удаляем вершину из стека вызовов
	recStack[v] = false
	return false
}

// Функция для проверки, можно ли топологически отсортировать граф
func isTopologicalSortPossible(graph map[string][]string) bool {
	// Отмечаем все вершины как непосещенные
	visited := make(map[string]bool)
	// Стек для отслеживания циклов
	recStack := make(map[string]bool)

	// Вызываем рекурсивную функцию для каждой непосещенной вершины
	for v := range graph {
		if !visited[v] {
			if isCyclicUtil(graph, v, visited, recStack) {
				return false
			}
		}
	}

	return true
}

func main() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {"D"},
		"D": {},
	}

	fmt.Println(isTopologicalSortPossible(graph)) // Вывод: true
}
