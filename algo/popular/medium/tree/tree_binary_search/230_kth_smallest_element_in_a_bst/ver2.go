package main

func kthSmallestV2(root *TreeNode, k int) int {

	target, _ := inorder(root, k, 0)
	return target
}

// target , idx
func inorder(root *TreeNode, k, idx int) (int, int) {

	if root == nil {
		return -1, idx
	}
	target := 0
	target, idx = inorder(root.Left, k, idx)

	if target != -1 {
		return target, idx
	}

	if idx+1 == k {
		return root.Val, idx
	}
	return inorder(root.Right, k, idx+1)
}
