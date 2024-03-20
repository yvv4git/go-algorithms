package main

func buildTreeV2(inorder []int, postorder []int) *TreeNode {
	/*
		METHOD: Iterative

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве.
		Это обусловлено тем, что мы проходим по каждому узлу в postorder обходе и для каждого узла выполняем постоянное количество операций,
		не зависящих от n.

		SPACE COMPLEXITY: O(n), так как мы храним позиции всех узлов в inorder обходе в словаре (map), что требует O(n) дополнительной памяти.
		Также мы используем дополнительную память для хранения узлов дерева, что также требует O(n) памяти.
		В итоге, общая пространственная сложность составляет O(n).
	*/
	// Если inorder пуст, дерево не может быть построено, поэтому возвращаем nil
	if len(inorder) == 0 {
		return nil
	}

	// Создаем словарь, где ключ - элемент из inorder, а значение - его индекс
	positions := make(map[int]int)
	for i, elem := range inorder {
		positions[elem] = i
	}

	// Последний элемент в postorder - это корень дерева
	root := &TreeNode{Val: postorder[len(postorder)-1]}

	// Проходим по всем элементам в postorder, кроме последнего (корня)
	for i := len(postorder) - 2; i >= 0; i-- {
		// Создаем новый узел с текущим элементом postorder
		node := &TreeNode{Val: postorder[i]}
		current := root

		// Ищем место для нового узла в дереве
		for {
			currentPosition := positions[current.Val]
			nodePosition := positions[node.Val]

			// Если позиция нового узла больше позиции текущего узла, ищем вправо
			if nodePosition > currentPosition {
				// Если правый потомок текущего узла не пустой, переходим к нему
				if current.Right != nil {
					current = current.Right
				} else {
					// Иначе добавляем новый узел как правый потомок текущего узла
					current.Right = node
					break
				}
			} else {
				// Если позиция нового узла меньше или равна позиции текущего узла, ищем влево
				if current.Left != nil {
					current = current.Left
				} else {
					// Иначе добавляем новый узел как левый потомок текущего узла
					current.Left = node
					break
				}
			}
		}
	}

	return root
}
