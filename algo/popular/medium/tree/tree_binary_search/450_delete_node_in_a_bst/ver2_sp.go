package main

func deleteNodeV2(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNodeV2(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNodeV2(root.Right, key)
	} else if root.Val == key {
		// leaf node
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			// has right node
			root.Val = findSuccessor(root)
			root.Right = deleteNodeV2(root.Right, root.Val)
		} else {
			// no right node but has left node
			root.Val = findPredecessor(root)
			root.Left = deleteNodeV2(root.Left, root.Val)
		}
	}
	return root
}

// Successor = "after node", i.e. the next node, or the smallest node after the current one.
func findSuccessor(node *TreeNode) int {
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}

// Predecessor = "before node", i.e. the previous node, or the largest node before the current one.
func findPredecessor(node *TreeNode) int {
	node = node.Left
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}
