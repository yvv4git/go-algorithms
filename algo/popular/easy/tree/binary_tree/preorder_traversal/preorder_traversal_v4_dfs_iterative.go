package preorder_traversal

func preorderTraversalV4(root *TreeNode) []int {
	/*
		Method: DFS iterative.
		Time complexity : O(n)
		Space complexity : O(n) - n used for stack.
	*/
	stack := []*TreeNode{root}
	result := []int{}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}

		result = append(result, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}

	return result
}
