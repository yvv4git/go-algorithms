package main

import "fmt"

// Функция для создания нового узла бинарного дерева
func newTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// Функция для построения бинарного дерева из предварительного и симметричного обходов
func buildTree(preorder []int, inorder []int) *TreeNode {
	/*
		METHOD: Iterative

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы обходим каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в стеке все узлы дерева.
	*/
	if len(preorder) == 0 {
		return nil
	}

	// Инициализация корня дерева
	root := newTreeNode(preorder[0])

	// Инициализация стека для отслеживания узлов
	stack := []*TreeNode{root}
	var inorderIndex int

	// Проход по предварительному обходу
	for i := 1; i < len(preorder); i++ {
		value := preorder[i]
		node := stack[len(stack)-1]

		if node.Val != inorder[inorderIndex] {
			node.Left = newTreeNode(value)
			stack = append(stack, node.Left)
		} else {
			for len(stack) != 0 && stack[len(stack)-1].Val == inorder[inorderIndex] {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				inorderIndex++
			}
			node.Right = newTreeNode(value)
			stack = append(stack, node.Right)
		}
	}

	return root
}

// Функция для печати дерева в симметричном порядке
func printInorder(node *TreeNode) {
	if node == nil {
		return
	}
	printInorder(node.Left)
	fmt.Printf("%d ", node.Val)
	printInorder(node.Right)
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)

	// Печать дерева в симметричном порядке для проверки
	printInorder(root)
}
