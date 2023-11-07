package _22_count_complete_tree_nodes

func countNodesV6(root *TreeNode) int {
	/*
		Method: DFS InorderTraversal.
		Time complexity : O(n)
		Space complexity : O(1).
	*/
	var cnt = 0

	inorderTraversal(root, &cnt)

	return cnt
}

func inorderTraversal(node *TreeNode, cnt *int) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left, cnt)
	*cnt++
	inorderTraversal(node.Right, cnt)
}
