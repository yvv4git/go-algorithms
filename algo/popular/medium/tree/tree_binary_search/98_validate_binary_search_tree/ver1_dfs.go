package main

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	/*
		METHOD: DFS
		Time complexity: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу.
		Space complexity: O(h), где h - высота дерева, так как в худшем случае глубина рекурсии может достигать h.
	*/
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	// Проверяем, что значение узла находится в допустимом диапазоне
	if node.Val <= min || node.Val >= max {
		return false
	}

	// Проверяем левое и правое поддеревья
	return validate(node.Left, min, node.Val) && validate(node.Right, node.Val, max)
}

func main() {
	// Создаем дерево для теста
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 8}

	// Проверяем дерево
	println(isValidBST(root)) // Должно вывести: true
}
