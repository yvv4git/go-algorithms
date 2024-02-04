package _971_find_if_path_exists_in_graph

import "container/list"

func validPathV2(n int, edges [][]int, start int, end int) bool {
	/*
		METHOD: BFS
		TIME COMPLEXITY: O(n + m), где n - количество вершин, m - количество ребер
		SPACE COMPLEXITY: O(n + m), где n - количество вершин, m - количество ребер

		Time complexity
		Временная сложность: O(n + m), где n - количество вершин, m - количество ребер.
		Это связано с тем, что в худшем случае мы посетим каждую вершину и каждое ребро.

		Space complexity
		Пространственная сложность: O(n + m), где n - количество вершин, m - количество ребер.
		Это связано с тем, что мы храним граф в виде списка смежности, где для каждой вершины мы храним список ее соседей.
	*/

	// Создаем граф, где индексы - вершины, а значения - списки смежных вершин
	graph := make([][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	// Создаем список посещенных вершин
	visited := make([]bool, n)

	// Создаем очередь для BFS
	queue := list.New()
	queue.PushBack(start)

	// Пока очередь не пуста
	for queue.Len() > 0 {
		// Извлекаем вершину из очереди
		node := queue.Front().Value.(int)
		queue.Remove(queue.Front())

		// Если это конечная вершина, то путь существует
		if node == end {
			return true
		}

		// Если вершина еще не посещена
		if !visited[node] {
			// Помечаем ее как посещенную
			visited[node] = true

			// Добавляем все ее соседей в очередь
			for _, neighbor := range graph[node] {
				queue.PushBack(neighbor)
			}
		}
	}

	// Если мы не нашли конечную вершину, то путь не существует
	return false
}
