package postorder_traversal

func postorderTraversalV2(root *TreeNode) []int {
	/*
		METHOD: DFS iterative.
		Time complexity : O(n)
		Space complexity : O(n) - n used for stack.
	*/
	stack := []*TreeNode{root}
	postorder := []int{}

	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		postorder = append([]int{node.Val}, postorder...)
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}

	return postorder
}
