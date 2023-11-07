package _22_count_complete_tree_nodes

func countNodesV5(root *TreeNode) int {
	/*
		Method: Binary search.
		Time complexity: O(logN * logN)
		Space complexity: O(1)
	*/
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	res := 1

	if leftHeight == rightHeight {
		//we go to right
		res += (1 << leftHeight) - 1 //left subtree nodes
		res += countNodesV5(root.Right)
	} else {
		res += (1 << rightHeight) - 1 //right subtree nodes
		res += countNodesV5(root.Left)
	}

	return res
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + getHeight(node.Left)
}
