package main

// Функция для вычисления среднего значения по уровням дерева с использованием DFS
func averageOfLevelsDFSV2(root *TreeNode) []float64 {
	/*
		METHOD: DFS
		В этом коде мы используем DFS для обхода дерева по уровням.
		Мы создаем два словаря: levelSums для хранения суммы значений узлов на каждом уровне и levelCounts для хранения количества узлов на каждом уровне.
		Затем мы вычисляем среднее значение на каждом уровне и возвращаем его в виде среза.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу только один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в стеке рекурсивных вызовов все узлы дерева.
	*/
	// Инициализация словаря для хранения суммы значений узлов и их количества на каждом уровне
	levelSums := make(map[int]float64)
	levelCounts := make(map[int]int)

	// Вызов вспомогательной функции для обхода дерева
	dfs(root, 0, levelSums, levelCounts)

	// Подсчет средних значений на каждом уровне
	result := make([]float64, len(levelSums))
	for level, sum := range levelSums {
		result[level] = sum / float64(levelCounts[level])
	}

	return result
}

// Вспомогательная функция для обхода дерева в глубину
func dfs(node *TreeNode, level int, levelSums map[int]float64, levelCounts map[int]int) {
	if node == nil {
		return
	}

	// Обновление суммы значений узлов и их количества на текущем уровне
	levelSums[level] += float64(node.Val)
	levelCounts[level]++

	// Рекурсивный вызов для левого и правого дочерних узлов
	dfs(node.Left, level+1, levelSums, levelCounts)
	dfs(node.Right, level+1, levelSums, levelCounts)
}

//func main() {
//	// Создание дерева для примера
//	root := &TreeNode{
//		Val: 3,
//		Left: &TreeNode{
//			Val: 9,
//		},
//		Right: &TreeNode{
//			Val: 20,
//			Left: &TreeNode{
//				Val: 15,
//			},
//			Right: &TreeNode{
//				Val: 7,
//			},
//		},
//	}
//
//	// Вызов функции и вывод результата
//	averages := averageOfLevelsDFS(root)
//	fmt.Println(averages) // Должно вывести: [3 14.5 11.5]
//}
