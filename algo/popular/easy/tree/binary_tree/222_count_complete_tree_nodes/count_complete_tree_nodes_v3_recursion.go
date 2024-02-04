package _22_count_complete_tree_nodes

func counter(count int, x *TreeNode) int {
	/*
		METHOD: Recursion.
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	if x == nil {
		return 0
	}

	count += counter(1, x.Left)
	count += counter(1, x.Right)

	return count
}

func countNodesV3(root *TreeNode) int {
	return counter(1, root)
}
