package main

import "math"

// Функция для поиска минимальной разницы между узлами в бинарном дереве поиска
func minDiffInBST(root *TreeNode) int {
	/*
		METHOD: DFS In-order
		TIME COMPLEXITY: O(n), где n - количество узлов в дереве. Это связано с тем, что мы посещаем каждый узел только один раз.
		SPACE COMPLEXITY: O(n), так как в худшем случае (когда дерево сбалансировано) мы можем иметь n рекурсивных вызовов функции inorder в стеке вызовов.

		In-order, pre-order, и post-order - это разные способы обхода дерева.
		1. In-order (in-order traversal) - это обход дерева в порядке лево-корень-право.
		Это означает, что мы сначала посещаем левое поддерево, затем корень, и, наконец, правое поддерево.

		2. Pre-order (pre-order traversal) - это обход дерева в порядке корень-лево-право.
		Это означает, что мы сначала посещаем корень, затем левое поддерево, и, наконец, правое поддерево.

		3. Post-order (post-order traversal) - это обход дерева в порядке лево-право-корень.
		Это означает, что мы сначала посещаем левое поддерево, затем правое поддерево, и, наконец, корень.

		DFS (Depth-First Search) - это алгоритм обхода или поиска графа или дерева.
		В DFS мы начинаем с корня (или произвольного узла) и посещаем каждый преемников узла до самого нижнего уровня.
		Если мы дошли до самого нижнего уровня, мы обходим уровень выше. DFS может быть реализован с помощью стека.

		In-order, pre-order, и post-order - это частные случаи DFS, где мы изменяем порядок посещения узлов.
	*/
	// Инициализация предыдущего узла и минимальной разницы
	prev := -1
	minDiff := math.MaxInt64

	// Функция для обхода дерева в порядке in-order (слева-корень-право)
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		// Базовый случай: если узел пустой, то выходим из рекурсии
		if node == nil {
			return
		}

		// Рекурсивный вызов для левого поддерева
		inorder(node.Left)

		// Если предыдущий узел не равен -1 и разница между текущим узлом и предыдущим меньше минимальной разницы,
		// то обновляем минимальную разницу
		if prev != -1 && node.Val-prev < minDiff {
			minDiff = node.Val - prev
		}

		// Обновляем предыдущий узел
		prev = node.Val

		// Рекурсивный вызов для правого поддерева
		inorder(node.Right)
	}

	// Запуск обхода in-order от корня дерева
	inorder(root)

	// Возвращаем минимальную разницу
	return minDiff
}
