package main

// TreeNode - tree node implementation
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBinaryTreeToArray(root *TreeNode) (result []int) {
	m := make(map[int]int)

	var processor func(*TreeNode, int)

	processor = func(node *TreeNode, idx int) {
		m[idx] = node.Val

		if node.Left != nil {
			idxLeft := (2 * idx) + 1
			processor(node.Left, idxLeft)
		}

		if node.Right != nil {
			idxRight := (2 * idx) + 2
			processor(node.Right, idxRight)
		}
	}

	processor(root, 0)

	maxKey := 0
	for key := range m {
		if maxKey < key {
			maxKey = key
		}
	}

	//result = make([]int, maxKey) // Выделяем память одним разом
	for i := 0; i <= maxKey; i++ {
		if val, ok := m[i]; ok {
			result = append(result, val)
		} else {
			result = append(result, 0)
		}
	}

	return result
}
