package main

// TreeNode - implementations of node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Задача заключается в том, чтобы определить - является ли это бинарное дерево сбалансированным.
Сбалансированное бинарное дерево - это бинарное дерево, длина поддеревьев которого отличается
максимум на 1.
В данном случае вопрос решается с помощью алгоритма DFS.
*/
func isBalanced(root *TreeNode) bool {
	return visit(root) != -1
}

func visit(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := visit(root.Left)
	if left == -1 {
		return -1
	}

	right := visit(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
