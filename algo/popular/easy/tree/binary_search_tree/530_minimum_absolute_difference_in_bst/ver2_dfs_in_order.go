package main

import "math"

func getMinimumDifferenceV2(root *TreeNode) int {
	/*
		METHOD: DFS in-order
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)

		Depth-First Search (DFS) - это метод обхода графа или дерева, который использует стек.
		В DFS мы начинаем с корня (или произвольного узла) и посещаем каждый дочерний узел,
		пока не достигнем узла без дочерних узлов. Затем мы возвращаемся назад и посещаем другие дочерние узлы.

		In-Order Traversal (обход в порядке возрастания) - это метод обхода дерева,
		при котором сначала обходится левое поддерево, затем корень, и, наконец, правое поддерево.
		Для бинарного дерева поиска этот метод позволяет получить элементы в отсортированном порядке.
	*/

	// Инициализируем предыдущее значение и минимальную разницу
	prev := -1
	minDiff := math.MaxInt64

	// Определяем вспомогательную функцию inorder для обхода дерева в порядке возрастания
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Обходим левое поддерево
		inorder(node.Left)

		// Если предыдущее значение не равно -1 и разница между текущим значением узла
		// и предыдущим значением меньше текущей минимальной разницы,
		// то обновляем минимальную разницу
		if prev != -1 && node.Val-prev < minDiff {
			minDiff = node.Val - prev
		}
		// Обновляем предыдущее значение
		prev = node.Val

		// Обходим правое поддерево
		inorder(node.Right)
	}

	// Запускаем обход в порядке возрастания
	inorder(root)

	// Возвращаем минимальную разницу
	return minDiff
}
