package main

import (
	"math"
)

// Алгоритм Флойда-Уоршелла для поиска кратчайших путей между всеми парами вершин в графе.
func FloydWarshall(n int, edges [][]int) ([][]int, bool) {
	// Инициализация матрицы расстояний.
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = math.MaxInt32 // Бесконечность
		}
		dist[i][i] = 0 // Расстояние до самой себя равно 0
	}

	// Заполнение матрицы расстояний весами ребер.
	for _, edge := range edges {
		dist[edge[0]][edge[1]] = 1 // Предполагаем, что все ребра имеют одинаковый вес 1
		dist[edge[1]][edge[0]] = 1 // Для неориентированного графа
	}

	// Запуск алгоритма Флойда-Уоршелла.
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt32 && dist[k][j] != math.MaxInt32 && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Проверка на наличие отрицательных циклов.
	for i := 0; i < n; i++ {
		if dist[i][i] < 0 {
			return nil, false // Обнаружен отрицательный цикл
		}
	}

	return dist, true
}

// ValidPathV3 использует алгоритм Флойда-Уоршелла для определения существования пути между двумя вершинами.
func validPathV3(n int, edges [][]int, source int, destination int) bool {
	/*
		METHOD: Floyd Warshall / Алгоритм Флойда-Уоршела

		TIME COMPLEXITY: O(V^3), где V - количество вершин в графе.
		Это делает его эффективным для графов с небольшим количеством вершин, но не очень эффективным для больших графов.

		SPACE COMPLEXITY: O(V^2), так как мы должны хранить матрицу расстояний для всех пар вершин.
	*/
	dist, noNegativeCycles := FloydWarshall(n, edges)
	if !noNegativeCycles {
		return false // Обнаружен отрицательный цикл
	}

	// Если расстояние от источника до пункта назначения не является бесконечностью, существует путь.
	return dist[source][destination] != math.MaxInt32
}

//func main() {
//	// Пример использования функции validPath
//	n := 4
//	edges := [][]int{{0, 1}, {1, 2}, {2, 0}, {3, 0}}
//	source := 0
//	destination := 2
//	pathExists := validPath(n, edges, source, destination)
//	if pathExists {
//		println("Путь существует между вершинами", source, "и", destination)
//	} else {
//		println("Путь не существует между вершинами", source, "и", destination)
//	}
//}
