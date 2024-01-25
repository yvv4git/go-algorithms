package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Postorder обход
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := make([]int, 0)
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}

func main() {
	/*
		В этом примере мы создаем бинарное дерево и вызываем функцию postorderTraversal для корня дерева.
		Результат будет список значений узлов в порядке, определяемом обходом в постфиксном порядке (левое поддерево, правое поддерево, корень).
	*/
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

	fmt.Println(postorderTraversal(root)) // Выведет: [4 5 2 3 1]
}
