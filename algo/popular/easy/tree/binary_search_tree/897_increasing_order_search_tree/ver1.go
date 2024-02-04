package main

func increasingBST(root *TreeNode) *TreeNode {
	/*
		METHOD: Dummy node
		TIME COMPLEXITY: O(n), где n - количество узлов в дереве. Это связано с тем, что мы посещаем каждый узел только один раз.
		Space complexity: O(1), так как мы не используем дополнительное пространство, кроме необходимого для рекурсивных вызовов.

		Метод:
		Метод, используемый для решения задачи, называется "In-Order Traversal" (обход в порядке "слева-корень-справа").
		Этот метод используется для обхода дерева в порядке возрастания значений.
		В этом случае мы используем его для обхода дерева в порядке возрастания значений,
		чтобы создать новое дерево, которое будет отсортировано.

		Почему используется этот метод:
		Метод "In-Order Traversal" является эффективным для этой задачи, так как он позволяет нам обойти дерево в порядке возрастания значений,
		что позволяет нам легко создать новое дерево, которое будет отсортировано.
	*/
	// Создание фиктивного узла, который будет использоваться для начала нового дерева
	dummy := &TreeNode{}
	// Инициализация предыдущего узла как фиктивного узла
	prev := dummy

	// Объявление функции inorder, которая будет использоваться для обхода дерева
	var inorder func(*TreeNode)

	// Определение функции inorder
	inorder = func(node *TreeNode) {
		// Если узел пуст, то выйти из функции
		if node == nil {
			return
		}

		// Обход левого поддерева
		inorder(node.Left)
		// Отсоединить левого потомка, так как он больше не нужен
		node.Left = nil
		// Подключить текущий узел к предыдущему узлу
		prev.Right = node
		// Переместить предыдущий узел на текущий узел
		prev = node
		// Обход правого поддерева
		inorder(node.Right)
	}

	// Запуск обхода inorder от корня дерева
	inorder(root)
	// Возвращение нового корня дерева, который находится справа от фиктивного узла
	return dummy.Right
}
