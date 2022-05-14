package main

func isSymmetric(root *TreeNode) bool {
	return chekSymmetric(root, root)
}

func chekSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val &&
		chekSymmetric(left.Right, right.Left) &&
		chekSymmetric(left.Left, right.Right)

}

// TreeNode - tree node implementation
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
