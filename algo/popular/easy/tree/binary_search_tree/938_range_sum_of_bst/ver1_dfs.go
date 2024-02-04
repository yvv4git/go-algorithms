package main

// Функция для вычисления суммы значений узлов в диапазоне от L до R
func rangeSumBST(root *TreeNode, low int, high int) int {
	/*
		METHOD: DFS
		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу.
		SPACE COMPLEXITY: O(h), где h - высота дерева, так как в худшем случае глубина рекурсии может достигать высоты дерева.
	*/
	// Базовый случай: если узел пустой, то сумма равна 0
	if root == nil {
		return 0
	}

	// Инициализируем сумму
	sum := 0

	// Если значение узла находится в диапазоне от low до high, то добавляем его значение к сумме
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	// Рекурсивно вызываем функцию для левого и правого поддеревьев
	if root.Val > low {
		sum += rangeSumBST(root.Left, low, high)
	}
	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	// Возвращаем сумму
	return sum
}
