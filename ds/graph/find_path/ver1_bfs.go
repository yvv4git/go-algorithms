package main

// Функция validPath проверяет существование пути между двумя вершинами в графе
func validPath(n int, edges [][]int, source int, destination int) bool {
	/*
		METHOD: BFS

		TIME COMPLEXITY: O(V + E), где V - количество вершин, E - количество ребер.
		Это связано с тем, что в худшем случае алгоритм посетит каждую вершину и каждое ребро.

		SPACE COMPLEXITY: O(V + E), так как в худшем случае мы будем хранить список смежности для каждой вершины и очередь для BFS.
	*/
	// Создаем словарь для хранения списка смежности
	res := make(map[int][]int)

	// Заполняем словарь списками смежности для каждой вершины
	for _, edge := range edges {
		u := edge[0]
		v := edge[1]
		res[edge[0]] = append(res[edge[0]], v)
		res[edge[1]] = append(res[edge[1]], u)
	}

	// Создаем словарь для отслеживания посещенных вершин
	visited := make(map[int]bool)

	// Создаем очередь для BFS
	queue := make([]int, 0)

	// Добавляем исходную вершину в очередь
	queue = append(queue, source)

	// Помечаем исходную вершину как посещенную
	visited[source] = true

	// Пока очередь не пуста
	for len(queue) > 0 {
		// Извлекаем первый элемент из очереди
		val := queue[0]
		queue = queue[1:]

		// Получаем список смежности для текущей вершины
		temp := res[val]

		// Проходим по всем соседним вершинам
		for i := 0; i < len(temp); i++ {
			vct := temp[i]

			// Если соседняя вершина еще не посещена
			if visited[vct] == false {
				// Помечаем ее как посещенную
				visited[vct] = true

				// Добавляем ее в очередь
				queue = append(queue, vct)
			}
		}

		// Если пункт назначения посещен, возвращаем true
		if visited[destination] == true {
			return true
		}
	}

	// Если после обхода всех вершин пункт назначения не посещен, возвращаем false
	return visited[destination]
}
