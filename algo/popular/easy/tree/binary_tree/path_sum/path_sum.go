package path_sum

func hasPathSum(root *TreeNode, targetSum int) bool {
	/*
		Time complexity: O(n)
		Space complexity: O(n)

		This is a depth-first search that decrements targetSum by the current node's value.
		If we reach a leaf node (left & right are nil) and target sum is zero, return true.
	*/
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
