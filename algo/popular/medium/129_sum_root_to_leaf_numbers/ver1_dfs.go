package main

import "fmt"

func sumNumbers(root *TreeNode) int {
	/*
		METHOD: DFS
		Для решения этой задачи мы будем использовать подход на основе рекурсии DFS.
		Мы будем обходить каждый узел дерева, передавая в рекурсивную функцию текущее значение числа, которое мы формируем.
		На каждом узле мы будем умножать текущее значение на 10 и прибавлять значение узла, чтобы получить новое число.
		Если мы дошли до листа, мы добавляем полученное число к сумме.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы обходим каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n) в худшем случае, когда дерево является сбалансированным,
		и O(log n) в среднем случае, когда дерево сбалансировано. Это происходит из-за использования рекурсии для обхода узлов дерева.
	*/
	return dfs(root, 0)
}

func dfs(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	// Обновляем текущую сумму, умножая ее на 10 и добавляя значение текущего узла
	currentSum = currentSum*10 + node.Val

	// Если это лист, возвращаем текущую сумму
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	// Рекурсивно вызываем функцию для левого и правого поддеревьев и возвращаем их сумму
	return dfs(node.Left, currentSum) + dfs(node.Right, currentSum)
}

func main() {
	// Пример использования
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	fmt.Println(sumNumbers(root)) // Вывод: 25
}
