package main

import "fmt"

func dfs(node *TreeNode, leaves *[]int) {
	if node != nil {
		if node.Left == nil && node.Right == nil {
			*leaves = append(*leaves, node.Val)
		}
		dfs(node.Left, leaves)
		dfs(node.Right, leaves)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	/*
		METHOD: DFS recursion
		Для решения этой задачи нам нужно будет реализовать две функции:
		1. Функцию dfs (глубокий поиск в глубину), которая будет обходить дерево и собирать листья в порядке обхода.
		2. Функцию leafSimilar, которая будет использовать dfs для двух деревьев и сравнивать их листья.

		Функция dfs будет рекурсивно обходить дерево, добавляя листья в список, когда они будут найдены.
		Функция leafSimilar будет использовать dfs для двух деревьев и сравнивать их листья.
		Если листья совпадают, функция вернет true, иначе - false.

		TIME COMPLEXITY: O(T1 + T2), где T1 и T2 - количество узлов в первом и втором дереве соответственно.
		Это обусловлено тем, что мы проходим по каждому дереву дважды: один раз для сбора листьев (где T - количество узлов в дереве),
		и второй раз для сравнения листьев.

		SPACE COMPLEXITY: O(T1 + T2), так как мы храним листья обоих деревьев в отдельных структурах данных.
		Каждая структура может содержать не более T1 и T2 элементов, если деревья содержат максимальное количество листьев.
		Поэтому общее количество элементов в структурах данных будет не больше 2T, где T - максимальное количество листьев в обоих деревьях.
	*/
	leaves1 := []int{}
	leaves2 := []int{}
	dfs(root1, &leaves1)
	dfs(root2, &leaves2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := range leaves1 {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}

func main() {
	// Пример использования
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 5}
	root1.Right = &TreeNode{Val: 1}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Right.Left = &TreeNode{Val: 9}
	root1.Right.Right = &TreeNode{Val: 8}

	root2 := &TreeNode{Val: 3}
	root2.Left = &TreeNode{Val: 5}
	root2.Right = &TreeNode{Val: 1}
	root2.Left.Left = &TreeNode{Val: 6}
	root2.Left.Right = &TreeNode{Val: 7}
	root2.Right.Left = &TreeNode{Val: 4}
	root2.Right.Right = &TreeNode{Val: 2}

	fmt.Println(leafSimilar(root1, root2)) // Вывод: false
}
