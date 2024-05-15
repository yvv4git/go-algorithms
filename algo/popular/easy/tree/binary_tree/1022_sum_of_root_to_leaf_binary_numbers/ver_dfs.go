package main

import "fmt"

// Функция для нахождения суммы путей от корня до листовых узлов.
func sumRootToLeaf(root *TreeNode) int {
	/*
		METHOD: DFS
		Мы будем обходить дерево в глубину, накапливая значение двоичного числа на каждом узле.
		Когда мы достигнем листового узла, мы добавим накопленное значение в сумму.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, потому что мы посещаем каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n) в худшем случае, когда дерево является сбалансированным, и O(log n) в среднем случае,
		когда дерево является сбалансированным. Это происходит из-за использования рекурсии для обхода дерева.
	*/
	return dfs(root, 0)
}

// Функция для обхода дерева в глубину.
func dfs(node *TreeNode, sum int) int {
	// Если узел пустой, возвращаем текущую сумму.
	if node == nil {
		return 0
	}

	// Обновляем текущее значение суммы.
	sum = sum<<1 | node.Val

	// Если это листовой узел, возвращаем текущую сумму.
	if node.Left == nil && node.Right == nil {
		return sum
	}

	// Рекурсивно обходим левое и правое поддеревья.
	return dfs(node.Left, sum) + dfs(node.Right, sum)
}

func main() {
	// Создаем дерево для примера.
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 1,
		},
	}

	// Находим сумму путей от корня до листовых узлов.
	sum := sumRootToLeaf(root)

	// Выводим результат.
	fmt.Println(sum) // Вывод: 22
}
