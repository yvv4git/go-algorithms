package main

// CloneGraphV3 копирует граф с использованием DFS с ограничением глубины.
func CloneGraphV3(node *Node) *Node {
	/*
		METHOD: ID BFS
		(Iterative Deepening Breadth-First Search) - это вариант BFS, где вместо последовательного обхода в ширину,
		мы используем BFS с различными начальными узлами, увеличивая глубину поиска по мере необходимости.
		Это может быть эффективным, если граф имеет широкие узлы, которые могут быть посещены до того,
		как будут обработаны узлы, находящиеся глубже.

		TIME COMPLEXITY: O(N + M), где N - количество узлов в графе, а M - количество рёбер.

		SPACE COMPLEXITY: O(N), так как в хэш-таблице visited для отслеживания уже скопированных узлов хранится один узел для каждого узла в исходном графе.
	*/
	if node == nil {
		return nil
	}

	// Хэш-таблица для отслеживания уже скопированных узлов.
	visited := make(map[*Node]*Node)

	// Рекурсивная функция для DFS с ограничением глубины.
	var clone func(node *Node, depth int) *Node
	clone = func(node *Node, depth int) *Node {
		if depth < 0 {
			return nil
		}

		// Если узел уже скопирован, возвращаем его копию.
		if v, ok := visited[node]; ok {
			return v
		}

		// Копируем узел.
		cloneNode := &Node{Val: node.Val}
		visited[node] = cloneNode

		// Копируем соседей.
		for _, neighbor := range node.Neighbors {
			cloneNeighbor := clone(neighbor, depth-1)
			if cloneNeighbor != nil {
				cloneNode.Neighbors = append(cloneNode.Neighbors, cloneNeighbor)
			}
		}

		return cloneNode
	}

	// Запускаем DFS с максимальной глубиной, например, 1000.
	return clone(node, 1000)
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
