package main

import "fmt"

func generateTrees(n int) []*TreeNode {
	/*
		METHOD: Recursion

		TIME COMPLEXITY: O(4^n / n^(3/2)), где n - количество узлов в дереве.
		Это связано с тем, что для каждого узла мы генерируем все возможные левое и правое поддеревья,
		и количество таких поддеревьев для каждого узла является каталаном, который растет экспоненциально с n.
		"Каталаном" в математике и компьютерной науке обычно относится к числу,
		которое представляет собой количество уникальных структур данных, которые могут быть построены из n элементов.
		Это число также называется числом Каталана.

		Например, число Каталана для n узлов в двоичном дереве поиска определяется формулой:
		C(n) = (2n)! / [(n + 1)! * n!]

		Это число дает количество уникальных двоичных деревьев поиска, которые можно построить из n различных ключей.

		SPACE COMPLEXITY: O(4^n / n^(3/2)), так как для хранения всех уникальных деревьев потребуется этот объем памяти.
		Каждое дерево требует O(n) памяти для хранения узлов, и существует O(4^n / n^(3/2)) деревьев.
	*/
	if n == 0 {
		return []*TreeNode{}
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	allTrees := []*TreeNode{}

	// Перебираем все возможные корни
	for i := start; i <= end; i++ {
		// Получаем все возможные левое поддерево с начала до i-1
		leftTrees := generate(start, i-1)

		// Получаем все возможные правое поддерево с i+1 до конца
		rightTrees := generate(i+1, end)

		// Объединяем все возможные левое и правое поддерево с корнем i
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currentTree := &TreeNode{Val: i}
				currentTree.Left = left
				currentTree.Right = right
				allTrees = append(allTrees, currentTree)
			}
		}
	}

	return allTrees
}

// Функция для печати дерева
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Print(root.Val, " ")
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	n := 3
	trees := generateTrees(n)

	for _, tree := range trees {
		printTree(tree)
		fmt.Println()
	}
}
