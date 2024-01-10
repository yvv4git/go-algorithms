package main

var res int

func amountOfTimeV1(root *TreeNode, start int) int {
	/*
		Метод: DFS (поиск в глубину)
		Временная сложность: O(n), где n - количество узлов в дереве, потому что мы проходим по каждому узлу ровно один раз.
		Пространственная сложность: O(n), так как в худшем случае (когда дерево сбалансировано) мы храним в стеке все узлы дерева.

		Сбалансированное дерево - это дерево, в котором для любого узла высота левого и правого поддеревьев отличается не более чем на 1.
		Это свойство делает сбалансированные деревья эффективными для многих операций, таких как поиск, вставка и удаление,
		которые обычно имеют сложность O(log n) в сбалансированном дереве.
	*/
	res = 0
	dfsSpecial(root, start)
	return res
}

func dfsSpecial(root *TreeNode, start int) (int, bool) {
	// Если узел пустой, то возвращаем 0 и false
	if root == nil {
		return 0, false
	}

	// Рекурсивно вызываем функцию для левого и правого поддеревьев
	left, foundInLeft := dfsSpecial(root.Left, start)
	right, foundInRight := dfsSpecial(root.Right, start)

	// Если значение текущего узла равно начальному значению, то обновляем res и возвращаем 1 и true
	if root.Val == start {
		res = max(res, max(left, right))
		return 1, true
	}

	// Если начальное значение найдено в левом поддереве, то обновляем res и возвращаем left + 1 и true
	if foundInLeft {
		res = max(res, right+left)
		return left + 1, true

		// Если начальное значение найдено в правом поддереве, то обновляем res и возвращаем right + 1 и true
	} else if foundInRight {
		res = max(res, right+left)
		return right + 1, true

		// Если начальное значение не найдено в текущем поддереве, то возвращаем максимальное время, необходимое для достижения узла start, и false
	} else {
		return max(left, right) + 1, false
	}
}
