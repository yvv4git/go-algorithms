package invert_binary_tree

func invertTreeV1(root *TreeNode) *TreeNode {
	/*
		METHOD: DFS, Recursion
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)
	*/
	if root != nil {
		root.Left, root.Right = invertTreeV1(root.Right), invertTreeV1(root.Left)
	}

	return root
}
