package main

// Функция amountOfTimeV2 вычисляет время, необходимое для достижения узла start в дереве root.
func amountOfTimeV2(root *TreeNode, start int) int {
	/*
		Метод: DFS (поиск в глубину)
		Временная сложность: O(n), где n - количество узлов в дереве, потому что мы проходим по каждому узлу ровно один раз.
		Пространственная сложность: O(n), так как в худшем случае (когда дерево сбалансировано) мы храним в стеке все узлы дерева.
	*/

	// Создаем пустую карту leafSize для хранения размеров листьев для каждого узла.
	leafSize := make(map[int][2]int)
	// Вычисляем размеры листьев для каждого узла в дереве.
	maxLeafSize(root, leafSize)
	// Начинаем обход дерева, начиная с корня, и ищем узел start.
	return traverseForFindTarget(root, leafSize, 0, start)
}

// Функция traverseForFindTarget выполняет обход дерева и ищет узел start.
func traverseForFindTarget(root *TreeNode, leafSize map[int][2]int, acc, target int) int {
	// Если узел пустой, возвращаем -1.
	if root == nil {
		return -1
	}
	// Если узел является целевым, возвращаем максимальный размер листа среди дочерних узлов целевого узла и текущего аккумулированного размера.
	if root.Val == target {
		return max(max(leafSize[root.Val][0], leafSize[root.Val][1]), acc)
	}
	// Если узел не является целевым, продолжаем обход левого и правого поддеревьев, обновляя аккумулированный размер.
	return max(
		traverseForFindTarget(root.Left, leafSize, max(acc, leafSize[root.Val][1])+1, target),
		traverseForFindTarget(root.Right, leafSize, max(acc, leafSize[root.Val][0])+1, target),
	)
}

// Функция maxLeafSize вычисляет размеры листьев для каждого узла в дереве.
func maxLeafSize(root *TreeNode, children map[int][2]int) int {
	// Если узел пустой, возвращаем 0.
	if root == nil {
		return 0
	}
	// Рекурсивно вычисляем размеры листьев для левого и правого поддеревьев.
	left := maxLeafSize(root.Left, children)
	right := maxLeafSize(root.Right, children)
	// Сохраняем размеры листьев для текущего узла.
	children[root.Val] = [2]int{left, right}
	// Возвращаем максимальный размер листа среди левого и правого поддеревьев, увеличенный на 1.
	return max(left, right) + 1
}
