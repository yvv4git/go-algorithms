package main

// CloneGraphV2 копирует граф с использованием BFS.
func CloneGraphV2(node *Node) *Node {
	/*
		METHOD: BFS

		TIME COMPLEXITY: O(N + M), где N - количество узлов в графе, а M - количество рёбер.
		Это связано с тем, что мы посещаем каждый узел и каждое ребро ровно один раз.

		SPACE COMPLEXITY: O(N), так как в хэш-таблице visited для отслеживания уже скопированных узлов хранится один узел для каждого узла в исходном графе.
		Очередь queue может содержать не более N элементов в худшем случае.
	*/
	if node == nil {
		return nil
	}

	// Хэш-таблица для отслеживания уже скопированных узлов.
	visited := make(map[*Node]*Node)

	// Очередь для BFS.
	queue := []*Node{node}

	// Копируем корневой узел.
	visited[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		// Извлекаем узел из очереди.
		current := queue[0]
		queue = queue[1:]

		// Копируем соседей текущего узла.
		for _, neighbor := range current.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				// Если сосед еще не скопирован, то копируем его и добавляем в очередь.
				visited[neighbor] = &Node{Val: neighbor.Val}
				queue = append(queue, neighbor)
			}
			// Добавляем копию соседа в список соседей копии текущего узла.
			visited[current].Neighbors = append(visited[current].Neighbors, visited[neighbor])
		}
	}

	return visited[node]
}

//func main() {
//	// Создаем тестовый граф.
//	node1 := &Node{Val: 1}
//	node2 := &Node{Val: 2}
//	node3 := &Node{Val: 3}
//	node4 := &Node{Val: 4}
//
//	node1.Neighbors = []*Node{node2, node4}
//	node2.Neighbors = []*Node{node1, node3}
//	node3.Neighbors = []*Node{node2, node4}
//	node4.Neighbors = []*Node{node1, node3}
//
//	// Копируем граф.
//	clonedGraph := CloneGraph(node1)
//
//	// Выводим значения узлов в скопированном графе.
//	fmt.Println(clonedGraph.Val)
//	for _, neighbor := range clonedGraph.Neighbors {
//		fmt.Println(neighbor.Val)
//	}
//}
