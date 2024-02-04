package main

import "fmt"

func main() {
	/*
		В этом примере мы создаем бинарное дерево с корнем 1, левого потомка 2 и правого потомка 3.
		У узла 2 есть два потомка: 4 и 5. Вызов inorderTraversal(root) вернет слайс [4, 2, 5, 1, 3],
		который является inorder обходом дерева.
	*/
	// Создаем бинарное дерево
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Вызываем функцию inorderTraversal
	result := inorderTraversal(root)

	// Выводим результат
	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	/*
		METHOD:
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность функции inorderTraversal в худшем случае составляет O(n), где n - количество узлов в дереве.
		Это происходит, когда дерево является сбалансированным, и мы посещаем каждый узел только один раз.

		Space complexity
		Пространственная сложность также составляет O(n) в худшем случае.
		Это происходит, когда дерево является сбалансированным, и мы храним в стеке все узлы дерева.


	*/
	// Если корень равен nil, то дерево пустое, и мы возвращаем пустой слайс.
	if root == nil {
		return []int{}
	}

	// Создаем слайс для хранения результата.
	result := make([]int, 0)

	// Рекурсивно вызываем функцию inorderTraversal для левого поддерева и добавляем результат в слайс.
	result = append(result, inorderTraversal(root.Left)...)

	// Добавляем значение корня в результат.
	result = append(result, root.Val)

	// Рекурсивно вызываем функцию inorderTraversal для правого поддерева и добавляем результат в слайс.
	result = append(result, inorderTraversal(root.Right)...)

	// Возвращаем результат.
	return result
}
