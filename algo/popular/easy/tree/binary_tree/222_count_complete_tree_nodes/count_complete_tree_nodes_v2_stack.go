package _22_count_complete_tree_nodes

func countNodesV2(root *TreeNode) int {
	/*
		METHOD: Stack.
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	var count int
	if root == nil {
		return count
	}

	count++
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// Вытаскиваем первый элемент.
		node := stack[len(stack)-1]  // Берем элемент сверху.
		stack = stack[:len(stack)-1] // Удаляем элемент сверху.

		if node.Left != nil {
			count++
			stack = append(stack, node.Left)
		}

		if node.Right != nil {
			count++
			stack = append(stack, node.Right)
		}
	}

	return count
}
