package main

func findTargetV2(root *TreeNode, k int) bool {
	/*
		Method: BFS
		Time complexity: O(n)
		Space complexity: O(n)

		BFS используется, когда нам нужно найти кратчайший путь от одной вершины до другой.
		Он проходит по всем вершинам, которые находятся на одном расстоянии от начальной вершины,
		прежде чем переходить к вершинам, которые находятся на расстоянии, равном длине пути от начальной вершины до текущей вершины + 1.

		DFS, напротив, используется, когда нам нужно проверить, существует ли путь между двумя вершинами,
		и если да, то найти этот путь. Он проходит по всем дочерним вершинам одной вершины,
		прежде чем переходить к дочерним вершинам другой вершины.

		В данном случае, нам нужно проверить, существует ли пара вершин, сумма значений которых равна заданному числу.
		BFS подходит лучше, так как он проходит по уровням дерева, начиная с корня, и проверяет сумму значений каждой пары вершин на каждом уровне.
		Если мы находим пару вершин, сумма значений которых равна заданному числу, то мы можем сказать, что существует такая пара вершин.
		Если мы не находим такую пару вершин после обхода всего дерева, то мы можем сказать, что такой пары вершин не существует.
	*/
	// Если корень дерева пуст, то возвращаем false
	if root == nil {
		return false
	}

	// Создаем очередь для обхода дерева в ширину
	queue := []*TreeNode{root}

	// Создаем множество для хранения значений узлов дерева
	set := make(map[int]bool)

	// Пока очередь не пуста
	for len(queue) > 0 {
		// Извлекаем первый элемент из очереди
		node := queue[0]
		queue = queue[1:]

		// Если в множестве есть значение, которое прибавлено к текущему значению узла дерева равно k,
		// то возвращаем true
		if set[k-node.Val] {
			return true
		}

		// Добавляем значение текущего узла в множество
		set[node.Val] = true

		// Если у текущего узла есть левый потомок, то добавляем его в очередь
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// Если у текущего узла есть правый потомок, то добавляем его в очередь
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	// Если не нашли таких значений, которые в сумме дают k, то возвращаем false
	return false
}