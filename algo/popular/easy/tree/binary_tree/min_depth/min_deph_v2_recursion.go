package min_depth

func minDepthV2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left != nil && root.Right != nil {
		return 1 + minimum(minDepthV2(root.Left), minDepthV2(root.Right))
	}

	if root.Left != nil {
		return 1 + minDepthV2(root.Left)
	} else {
		return 1 + minDepthV2(root.Right)
	}
}

func minimum(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
