package min_depth

type NodeWithLevel struct {
	*TreeNode
	Level int
}

func minDepthV4(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*NodeWithLevel{
		{
			TreeNode: root,
			Level:    1,
		},
	}
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Left == nil && node.Right == nil {
			return node.Level
		}
		if node.Left != nil {
			queue = append(queue, &NodeWithLevel{
				TreeNode: node.Left,
				Level:    node.Level + 1,
			})
		}
		if node.Right != nil {
			queue = append(queue, &NodeWithLevel{
				TreeNode: node.Right,
				Level:    node.Level + 1,
			})
		}
	}
	return -1
}
