package main

import "fmt"

// Определение структуры TreeNode для узла дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция для обхода дерева в змейку
func zigzagLevelOrder(root *TreeNode) [][]int {
	/*
		METHOD: BFS
		Для решения этой задачи мы можем использовать алгоритм обхода в ширину (BFS),
		так как нам нужно обойти дерево уровня за уровнем. В этом алгоритме мы используем очередь для хранения узлов на каждом уровне дерева.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы посещаем каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в очереди все узлы последнего уровня дерева.
	*/
	if root == nil {
		return [][]int{}
	}

	// Инициализация переменных
	result := [][]int{}
	queue := []*TreeNode{root}
	direction := 1 // Направление обхода: 1 - слева направо, -1 - справа налево

	// Пока очередь не пуста
	for len(queue) > 0 {
		level := []int{}        // Результат на текущем уровне
		levelSize := len(queue) // Размер очереди на текущем уровне

		// Обработка всех узлов на текущем уровне
		for i := 0; i < levelSize; i++ {
			// Извлечение первого элемента из очереди
			node := queue[0]
			queue = queue[1:]

			// Добавление значения узла в результат текущего уровня
			level = append(level, node.Val)

			// Добавление дочерних узлов в очередь
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Применение змейки, если надо
		if direction == -1 {
			reverse(level)
		}

		// Добавление результата текущего уровня в общий результат
		result = append(result, level)

		// Изменение направления обхода
		direction *= -1
	}

	return result
}

// Функция для разворота слайса
func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	// Пример использования
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

	result := zigzagLevelOrder(root)
	fmt.Println(result) // Должен вывести [[3], [20, 9], [15, 7]]
}
