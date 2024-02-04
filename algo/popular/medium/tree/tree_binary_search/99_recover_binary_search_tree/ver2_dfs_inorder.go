package main

func recoverTreeV2(root *TreeNode) {
	/*
		METHOD: DFS
		Time complexity: O(n)
		Space complexity: O(1)

		Метод решения задачи:
		1. Инициализируем два указателя на предыдущий и текущий узлы.
		2. Запускаем рекурсивный обход дерева (например, In-Order обход).
		3. В процессе обхода, мы сравниваем текущий узел с предыдущим узлом.
		Если текущий узел меньше предыдущего узла, то мы нашли два узла, которые были помечены неправильно.
		4. Мы также должны отслеживать первый узел, который мы нашли, так как он может быть первым узлом, который нарушает инвариант BST.
		5. После обхода, мы должны поменять значения двух узлов, которые были помечены неправильно.
	*/

	// Объявляем переменные prev, first, second, которые будут использоваться для отслеживания узлов, которые нарушают порядок возрастания.
	var prev, first, second *TreeNode

	// Объявляем функцию dfs, которая будет выполнять обход дерева в глубину.
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		// Если узел пустой, то мы дошли до листа дерева и возвращаемся назад.
		if node == nil {
			return
		}

		// Рекурсивный вызов функции для левого поддерева.
		dfs(node.Left)

		// Если предыдущий узел не пустой, то мы проверяем, нарушает ли текущий узел порядок возрастания.
		if prev != nil {
			// Если первый узел еще не был найден и значение предыдущего узла больше или равно значению текущего узла,
			// то мы запоминаем предыдущий узел как первый.
			if first == nil && prev.Val >= node.Val {
				first = prev
			}
			// Если первый узел уже был найден и значение предыдущего узла больше или равно значению текущего узла,
			// то мы запоминаем текущий узел как второй.
			if first != nil && prev.Val >= node.Val {
				second = node
			}
		}

		// После того, как мы обработали текущий узел, мы запоминаем его как предыдущий.
		prev = node

		// Рекурсивный вызов функции для правого поддерева.
		dfs(node.Right)
	}

	// Вызываем функцию dfs для корня дерева.
	dfs(root)

	// Меняем местами значения первого и второго узлов.
	first.Val, second.Val = second.Val, first.Val
}
