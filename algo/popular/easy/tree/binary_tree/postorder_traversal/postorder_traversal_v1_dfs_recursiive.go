package postorder_traversal

func postorderTraversalV1(root *TreeNode) []int {
	/*
		Method: DFS recursion.
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	if root == nil {
		return nil
	}

	left := postorderTraversalV1(root.Left)
	right := postorderTraversalV1(root.Right)

	return append(left, append(right, root.Val)...)
}
