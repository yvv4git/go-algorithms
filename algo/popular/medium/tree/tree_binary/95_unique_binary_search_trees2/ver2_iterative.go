package main

func generateTreesV2(n int) []*TreeNode {
	/*
		METHOD: Iterative

		TIME COMPLEXITY: O(4^n / n^(1/2)), где n - количество узлов в дереве.
		Это связано с тем, что для каждого количества узлов мы перебираем все возможные корни,
		и для каждого корня мы генерируем все возможные левое и правое поддеревья, которые уже были рассчитаны ранее.

		SPACE COMPLEXITY: O(4^n / n^(1/2)), так как для хранения всех уникальных деревьев потребуется этот объем памяти.
		Каждое дерево требует O(n) памяти для хранения узлов, и существует O(4^n / n^(1/2)) деревьев.
	*/
	if n == 0 {
		return []*TreeNode{}
	}

	// Инициализируем массив dp
	dp := make([][]*TreeNode, n+1)
	for i := range dp {
		dp[i] = []*TreeNode{}
	}

	// Базовый случай: дерево с 0 узлами
	dp[0] = append(dp[0], nil)

	// Заполняем dp для количества узлов от 1 до n
	for nodes := 1; nodes <= n; nodes++ {
		// Для каждого количества узлов, перебираем все возможные корни
		for root := 1; root <= nodes; root++ {
			left := root - 1
			right := nodes - root

			// Получаем все возможные левое и правое поддеревья
			for _, leftTree := range dp[left] {
				for _, rightTree := range dp[right] {
					// Создаем корень с текущим значением
					rootNode := &TreeNode{Val: root}

					// Присоединяем левое и правое поддеревья к корню
					rootNode.Left = cloneTree(leftTree, 0)
					rootNode.Right = cloneTree(rightTree, root)

					// Добавляем новое дерево в dp
					dp[nodes] = append(dp[nodes], rootNode)
				}
			}
		}
	}

	return dp[n]
}

// Функция для клонирования дерева с увеличением значений узлов на offset
func cloneTree(node *TreeNode, offset int) *TreeNode {
	if node == nil {
		return nil
	}
	newNode := &TreeNode{Val: node.Val + offset}
	newNode.Left = cloneTree(node.Left, offset)
	newNode.Right = cloneTree(node.Right, offset)
	return newNode
}
