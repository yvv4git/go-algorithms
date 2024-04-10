package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(root *TreeNode) int {
	/*
		METHOD: Dynamic programming
		В этом подходе мы будем использовать рекурсию, но будем сохранять уже вычисленные значения в хэш-таблице,
		чтобы избежать повторных вычислений.

		TIME COMPLEXITY: O(n), где n - количество домов в дереве, потому что мы посещаем каждый дом только один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в хэш-таблице значение для каждого дома.
	*/
	// Используем хэш-таблицу для хранения уже вычисленных значений
	memo := make(map[*TreeNode]int)
	return robHelper(root, memo)
}

func robHelper(root *TreeNode, memo map[*TreeNode]int) int {
	// Базовый случай: если дом не существует, возвращаем 0
	if root == nil {
		return 0
	}

	// Проверяем, есть ли уже вычисленное значение для этого дома
	if val, ok := memo[root]; ok {
		return val
	}

	// Если дом не был посещен ранее, вычисляем его максимальную сумму
	// Если мы воруем этот дом, то мы не можем ворваться в соседние дома
	// и получаем сумму, равную сумме воруемых домов в левом и правом поддереве
	val := root.Val
	if root.Left != nil {
		val += robHelper(root.Left.Left, memo) + robHelper(root.Left.Right, memo)
	}
	if root.Right != nil {
		val += robHelper(root.Right.Left, memo) + robHelper(root.Right.Right, memo)
	}

	// Если мы не воруем этот дом, то максимальная сумма равна
	// максимальной сумме воруемых домов в левом и правом поддереве
	val = max(val, robHelper(root.Left, memo)+robHelper(root.Right, memo))

	// Сохраняем вычисленное значение в хэш-таблице
	memo[root] = val

	return val
}

func main() {
	// Пример использования функции rob
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 1,
			},
		},
	}
	fmt.Println(rob(root)) // Вывод: 7
}
