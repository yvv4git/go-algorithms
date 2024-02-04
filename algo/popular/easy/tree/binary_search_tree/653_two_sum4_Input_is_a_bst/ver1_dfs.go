package main

func findTarget(root *TreeNode, k int) bool {
	/*
		METHOD: DFS
		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, поскольку мы посещаем каждый узел ровно один раз.
		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем поместить все узлы в множество.
	*/
	set := make(map[int]bool)
	return dfs(root, k, set)
}

// Функция dfs выполняет Depth-First Search (DFS) для дерева.
// Она проверяет, существует ли пара узлов, сумма которых равна k.
// Если существует, то возвращает true, иначе false.
func dfs(root *TreeNode, k int, set map[int]bool) bool {
	// Если узел пустой, то возвращаем false.
	if root == nil {
		return false
	}

	// Если в множестве set есть значение, которое прибавлением root.Val дает k,
	// то возвращаем true, так как мы нашли пару узлов, сумма которых равна k.
	if set[k-root.Val] {
		return true
	}

	// Добавляем значение текущего узла в множество set.
	set[root.Val] = true

	// Выполняем DFS для левого и правого поддеревьев.
	// Если хотя бы для одного из поддеревьев мы нашли пару узлов, сумма которых равна k,
	// то возвращаем true.
	return dfs(root.Left, k, set) || dfs(root.Right, k, set)
}
