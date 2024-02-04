package invert_binary_tree

func invertTree(root *TreeNode) *TreeNode {
	/*
		METHOD: Loop by Stack
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)
	*/
	if root == nil {
		return root
	}

	// Set up stack of nodes for iterative pre-order traversal
	nodeStack := []*TreeNode{}
	nodeStack = append(nodeStack, root)

	// Loop until all nodes have been hit
	for len(nodeStack) > 0 {

		// Pop one node off stack
		currentNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]

		// Push right node then left node to stack so we traverse left first when popping
		if currentNode.Right != nil {
			nodeStack = append(nodeStack, currentNode.Right)
		}
		if currentNode.Left != nil {
			nodeStack = append(nodeStack, currentNode.Left)
		}

		// Swap positions of current node's left and right branches
		currentNode.Left, currentNode.Right = currentNode.Right, currentNode.Left
	}

	return root
}
