package main

import "fmt"

// Функция для вычисления среднего значения по уровням дерева
func averageOfLevels(root *TreeNode) []float64 {
	/*
		METHOD: BFS
		В этом коде мы используем обход в ширину для обхода дерева по уровням.
		Мы создаем очередь, в которую мы помещаем узлы дерева, и извлекаем их по одному за раз.
		При этом мы также подсчитываем сумму значений узлов на каждом уровне и количество узлов на этом уровне.
		После того, как мы обошли все узлы на одном уровне, мы добавляем среднее значение на этом уровне в результирующий массив.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу только один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в очереди все узлы одного уровня.
	*/
	// Проверка на случай, если дерево пустое
	if root == nil {
		return []float64{}
	}

	// Инициализация переменных
	var result []float64
	queue := []*TreeNode{root}

	// Пока очередь не пуста
	for len(queue) > 0 {
		// Подсчет количества узлов на текущем уровне
		levelSize := len(queue)
		levelSum := 0

		// Проход по всем узлам на текущем уровне
		for i := 0; i < levelSize; i++ {
			// Извлечение узла из очереди
			node := queue[0]
			queue = queue[1:]

			// Добавление значения узла к сумме на текущем уровне
			levelSum += node.Val

			// Добавление дочерних узлов в очередь
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Вычисление среднего значения на текущем уровне и добавление его в результат
		result = append(result, float64(levelSum)/float64(levelSize))
	}

	return result
}

func main() {
	// Создание дерева для примера
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	// Вызов функции и вывод результата
	averages := averageOfLevels(root)
	fmt.Println(averages) // Должно вывести: [3 14.5 11.5]
}
