package main

import "math"

// Функция getMinimumDifference находит минимальную разницу между значениями узлов в бинарном дереве.
func getMinimumDifference(root *TreeNode) int {
	/*
		Method: DFS in-order
		Time complexity: O(n), так как в худшем случае функция может хранить все значения узлов в дереве в срезе.
		Space complexity: O(n), так как в худшем случае функция может хранить все значения узлов в дереве в срезе.

		В данном коде используется In-Order Depth-First Search (DFS), так как мы сначала обходим левое поддерево,
		затем обрабатываем узел, а затем обходим правое поддерево.
	*/
	// Создаем пустой срез для хранения значений узлов.
	values := make([]int, 0, 10000) // 10^4 из условия задачи

	// Объявляем функцию traverse для обхода дерева в глубину.
	var traverse func(*TreeNode)
	traverse = func(root *TreeNode) {
		// Если узел пустой, то выходим из функции.
		if root == nil {
			return
		}

		// Значит использует метод обхода In-order(left, node, right).
		// Сначала обходим левое поддерево.
		traverse(root.Left)
		// Затем добавляем значение текущего узла в срез.
		values = append(values, root.Val)
		// После этого обходим правое поддерево.
		traverse(root.Right)
	}

	// Запускаем обход дерева.
	traverse(root)

	// Инициализируем переменную result максимальным значением int32.
	result := math.MaxInt32
	// Проходим по срезу значений и находим минимальную разницу между соседними значениями.
	for i := 1; i < len(values); i++ {
		result = min(result, values[i]-values[i-1])
	}

	// Возвращаем минимальную найденную разницу.
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
