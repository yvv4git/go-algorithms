package invert_binary_tree

func invertTreeV1(root *TreeNode) *TreeNode {
	/*
		METHOD: DFS, Recursion
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	if root != nil {
		root.Left, root.Right = invertTreeV1(root.Right), invertTreeV1(root.Left)
	}

	return root
}
