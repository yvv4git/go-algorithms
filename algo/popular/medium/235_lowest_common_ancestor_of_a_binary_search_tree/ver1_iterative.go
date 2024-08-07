package main

// Функция для нахождения наименьшего общего предка (LCA) в бинарном дереве поиска
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	/*
		METHOD: Iterative
		Этот алгоритм не использует стек или очередь, которые обычно применяются в DFS и BFS, соответственно.
		Вместо этого он использует итеративный подход с одной переменной для хранения текущего узла, что делает его очень эффективным для BST.

		Алгоритм работает следующим образом:
		1. Начиная с корня, он сравнивает значения узлов p и q с текущим узлом.
		2. Если оба значения меньше текущего узла, он переходит к левому поддереву.
		3. Если оба значения больше текущего узла, он переходит к правому поддереву.
		4. Если одно значение меньше, а другое больше текущего узла, то текущий узел является наименьшим общим предком.

		TIME COMPLEXITY: O(h), где h — высота дерева. Это связано с тем, что в худшем случае мы пройдем от корня до одного из листьев дерева.

		SPACE COMPLEXITY: O(1), так как мы не используем дополнительные структуры данных, а только переменные для хранения текущего узла.
	*/
	// Начинаем с корня дерева
	current := root

	for current != nil {
		// Если оба значения p и q меньше текущего узла, идем влево
		if p.Val < current.Val && q.Val < current.Val {
			current = current.Left
		} else if p.Val > current.Val && q.Val > current.Val {
			// Если оба значения p и q больше текущего узла, идем вправо
			current = current.Right
		} else {
			// Если одно значение меньше, а другое больше текущего узла, то текущий узел и есть LCA
			return current
		}
	}
	return nil
}
