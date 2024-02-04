package main

import "fmt"

func main() {
	/*
		В этом примере мы создаем бинарное дерево с корнем 1, левого потомка 2 и правого потомка 3. У узла 2 есть два потомка: 4 и 5.
		Вызов preorderTraversal(root) вернет слайс [1, 2, 4, 5, 3], который является preorder обходом дерева.
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

	// Вызываем функцию preorderTraversal
	result := preorderTraversal(root)

	// Выводим результат
	fmt.Println(result) // [1 2 4 5 3]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Preorder обход
func preorderTraversal(root *TreeNode) []int {
	/*
		METHOD:
		TIME COMPLEXITY: O(n)
		Space complexity: O(1) или O(n)

		Time complexity
		Временная сложность функции preorderTraversal в худшем случае составляет O(n), где n - количество узлов в дереве.
		Это происходит, когда дерево является сбалансированным, и мы посещаем каждый узел только один раз.

		Space complexity
		Пространственная сложность также составляет O(n) в худшем случае. Это происходит,
		когда дерево является сбалансированным, и мы храним в стеке все узлы дерева.
	*/

	// Если корень равен nil, то дерево пустое, и мы возвращаем пустой слайс.
	if root == nil {
		return []int{}
	}

	// Создаем слайс для хранения результата.
	result := make([]int, 0)

	// Добавляем значение корня в результат.
	result = append(result, root.Val)

	// Рекурсивно вызываем функцию preorderTraversal для левого поддерева и добавляем результат в слайс.
	result = append(result, preorderTraversal(root.Left)...)

	// Рекурсивно вызываем функцию preorderTraversal для правого поддерева и добавляем результат в слайс.
	result = append(result, preorderTraversal(root.Right)...)

	// Возвращаем результат.
	return result
}
