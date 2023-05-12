package min_depth

func minDepthV1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode
	enqueue(&queue, root)

	var res int

	for len(queue) > 0 {
		queueSize := len(queue)

		for i := 0; i < queueSize; i++ {
			dequeuedItem := dequeue(&queue)

			if i == 0 {
				res++
			}

			if dequeuedItem.Left == nil && dequeuedItem.Right == nil {
				return res
			}

			if dequeuedItem.Left != nil {
				enqueue(&queue, dequeuedItem.Left)
			}

			if dequeuedItem.Right != nil {
				enqueue(&queue, dequeuedItem.Right)
			}
		}
	}

	return res
}

func enqueue(queue *[]*TreeNode, item *TreeNode) {
	if queue == nil {
		panic("nil pointer")
	}

	*queue = append(*queue, item)
}

func dequeue(queue *[]*TreeNode) *TreeNode {
	if queue == nil {
		panic("nil pointer")
	}

	if len(*queue) == 0 {
		panic("empty queue")
	}

	dequeuedItem := (*queue)[0]
	*queue = (*queue)[1:]

	return dequeuedItem
}
