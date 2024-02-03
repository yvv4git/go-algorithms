package main

func kthSmallest(root *TreeNode, k int) int {
	/*
		Method: In-Order Traversal (Обход дерева в порядке возрастания)
		Time complexity: O(H+k), где H - высота дерева, а k - номер элемента, который мы ищем.
		Space complexity: O(H), где H - высота дерева.

		Обход дерева в порядке возрастания работает следующим образом:
		1. Сначала обходим левое поддерево.
		2. Затем обрабатываем корень.
		3. После этого обходим правое поддерево.
	*/

	// Инициализируем стек для обхода дерева в порядке возрастания
	stack := []*TreeNode{}

	// Начинаем с корня
	for {
		// Пока мы можем идти влево, добавляем все узлы в стек
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// Берем последний узел из стека
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Уменьшаем k
		k--

		// Если k становится равным 0, мы нашли k-й наименьший элемент
		if k == 0 {
			return root.Val
		}

		// Переходим вправо
		root = root.Right
	}
}
