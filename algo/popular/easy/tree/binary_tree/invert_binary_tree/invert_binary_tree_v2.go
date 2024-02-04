package invert_binary_tree

func invertTreeV2(root *TreeNode) *TreeNode {
	/*
		METHOD: Recursion, DSF
		Time complexity: O(n)
		Space complexity: O(n)
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
