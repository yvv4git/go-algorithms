package main

// isSameTree - офигенный рекурсивный вариант решения проблемы
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// TreeNode - tree node implementation
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
