package main

import "fmt"

// Функция для преобразования дерева в связный список
func flatten(root *TreeNode) {
	/*
		METHOD: Recursion DFS
		Это алгоритм, который использует рекурсию для обхода дерева в порядке лево-корень-право.
		В процессе обхода мы будем изменять структуру дерева, чтобы каждый узел стал правым потомком предыдущего узла.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем потребовать дополнительное пространство для рекурсивного стека, когда дерево является сбалансированным.
	*/
	// Базовый случай: если узел пустой, то преобразовывать нечего
	if root == nil {
		return
	}

	// Рекурсивно преобразовать левое и правое поддеревья
	flatten(root.Left)
	flatten(root.Right)

	// Сохранить правое поддерево
	right := root.Right

	// Сделать левое поддерево корнем правого поддерева
	root.Right = root.Left
	root.Left = nil

	// Переместить правое поддерево в конец нового правого поддерева
	for root.Right != nil {
		root = root.Right
	}
	root.Right = right
}

// Функция для печати связного списка
func printList(root *TreeNode) {
	for root != nil {
		fmt.Println(root.Val)
		root = root.Right
	}
}

func main() {
	// Создание дерева для примера
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	// Преобразование дерева в связный список
	flatten(root)

	// Печать связного списка
	printList(root)
}
