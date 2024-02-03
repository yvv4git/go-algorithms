package main

// Функция amountOfTimeV3 вычисляет максимальное расстояние от узла start до любого другого узла в дереве.
func amountOfTimeV3(root *TreeNode, start int) int {
	/*
		Methods: BFS + backtracking
		Time complexity: O(n^2), в среднем случае, так как она вызывает функции bfs и backtracking.
		Space complexity: O(n)
	*/

	// Инициализация максимального расстояния и двух структур данных: nodes и footprint.
	// nodes - это словарь, где ключ - это значение узла, а значение - это список смежных узлов.
	// footprint - это словарь, где ключ - это значение узла, а значение - это пустая структура.
	// Это используется для отслеживания посещенных узлов.
	maxDistance := 0
	nodes := make(map[int][]int)
	footprint := make(map[int]struct{})
	footprint[start] = struct{}{}

	// Выполняет обход в ширину (BFS) дерева и заполняет словарь nodes.
	bfs(root, nodes)

	// Выполняет обход в глубину (DFS) с учетом словаря nodes и footprint,
	// чтобы найти максимальное расстояние от узла start до любого другого узла.
	backtracking(nodes, footprint, start, 0, &maxDistance)

	// Возвращает максимальное расстояние.
	return maxDistance
}

// Функция backtracking выполняет обход в глубину (DFS) с учетом словаря nodes и footprint.
// O(n^2) в худшем случае, когда граф не является деревом и содержит циклы. В среднем случае, когда граф является деревом, временная сложность составляет O(n^2).
func backtracking(nodes map[int][]int, footprint map[int]struct{}, start, currentDistance int, maxDistance *int) {
	// Получает список смежных узлов для текущего узла.
	edges, ok := nodes[start]
	if !ok {
		return
	}

	// Для каждого смежного узла, если он еще не посещен, то посещает его,
	// обновляет текущее расстояние и вызывает backtracking для него.
	// После этого, удаляет узел из footprint.
	for _, edge := range edges {
		if _, ok := footprint[edge]; ok {
			continue
		}
		footprint[edge] = struct{}{}
		backtracking(nodes, footprint, edge, currentDistance+1, maxDistance)
		delete(footprint, edge)
	}

	// Обновляет максимальное расстояние, если текущее расстояние больше.
	*maxDistance = max(currentDistance, *maxDistance)
}

// Функция bfs выполняет обход в ширину (BFS) дерева и заполняет словарь nodes.
func bfs(root *TreeNode, nodes map[int][]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		nodes[root.Val] = append(nodes[root.Val], root.Left.Val)
		nodes[root.Left.Val] = append(nodes[root.Left.Val], root.Val)
		bfs(root.Left, nodes)
	}

	if root.Right != nil {
		nodes[root.Val] = append(nodes[root.Val], root.Right.Val)
		nodes[root.Right.Val] = append(nodes[root.Right.Val], root.Val)
		bfs(root.Right, nodes)
	}
}
