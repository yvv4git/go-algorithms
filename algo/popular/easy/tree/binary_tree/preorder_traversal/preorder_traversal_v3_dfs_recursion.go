package preorder_traversal

func preorderTraversalV3(root *TreeNode) []int {
	/*
		Method: DFS recursion.
		Time complexity : O(n)
		Space complexity : O(1)
	*/

	if root == nil {
		return nil
	}

	result := []int{root.Val}
	left := preorderTraversalV3(root.Left)
	right := preorderTraversalV3(root.Right)

	return append(result, append(left, right...)...)
}
