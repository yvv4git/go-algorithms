package invert_binary_tree

func invertTreeV2(root *TreeNode) *TreeNode {
	/*
		METHOD: Recursion, DSF
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	if root == nil {
		return nil
	}

	return &TreeNode{
		Val:   root.Val,
		Left:  invertTreeV2(root.Right),
		Right: invertTreeV2(root.Left),
	}
}
