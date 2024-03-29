package main

import "fmt"

// Функция для обхода дерева в глубину и получения списка узлов, видных справа.
func rightSideView(root *TreeNode) []int {
	/*
		METHOD: DFS
		Мы будем обходить дерево, начиная с корня, и для каждого уровня будем сохранять только последний посещенный узел.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, потому что мы обходим каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в очереди все узлы одного уровня дерева.
	*/
	// Инициализация результирующего слайса.
	var result []int

	// Вызов вспомогательной функции для обхода дерева.
	dfs(root, &result, 0)

	// Возврат результата.
	return result
}

// Вспомогательная функция для обхода дерева в глубину.
func dfs(node *TreeNode, result *[]int, level int) {
	// Базовый случай: если узел пустой, то прекращаем выполнение.
	if node == nil {
		return
	}

	// Если это новый уровень, добавляем его в результат.
	if level == len(*result) {
		*result = append(*result, node.Val)
	}

	// Обновляем значение узла на этом уровне.
	(*result)[level] = node.Val

	// Обходим левое поддерево.
	dfs(node.Left, result, level+1)

	// Обходим правое поддерево.
	dfs(node.Right, result, level+1)
}

func main() {
	// Создание бинарного дерева для примера.
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	// Получение списка узлов, видных справа.
	view := rightSideView(root)

	// Вывод результата.
	fmt.Println(view) // Должно вывести: [1, 3, 4]
}
