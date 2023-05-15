package preorder_traversal

func preorderTraversalV2(root *TreeNode) []int {
	/*
		Method: BFS.
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	var res []int

	if root == nil {
		return res
	}
	var stack []*TreeNode
	push(&stack, root)

	for len(stack) > 0 {
		popedElement := pop(&stack)
		res = append(res, popedElement.Val)

		if popedElement.Right != nil {
			push(&stack, popedElement.Right)
		}

		if popedElement.Left != nil {
			push(&stack, popedElement.Left)
		}
	}

	return res
}

func push(stack *[]*TreeNode, item *TreeNode) {
	if stack == nil {
		panic("nil pointer")
	}

	*stack = append(*stack, item)
}

func pop(stack *[]*TreeNode) *TreeNode {
	if stack == nil {
		panic("nil pointer")
	}

	if len(*stack) == 0 {
		panic("empty stack")
	}

	popedElement := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1]

	return popedElement
}
