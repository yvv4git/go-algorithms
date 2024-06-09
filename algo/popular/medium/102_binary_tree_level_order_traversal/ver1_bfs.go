package main

import "fmt"

// Функция для обхода дерева в порядке уровней
func levelOrder(root *TreeNode) [][]int {
	/*
		METHOD: BFS
		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по каждому узлу ровно один раз.
		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в очереди все узлы одного уровня дерева.
	*/
	// Если корень nil, возвращаем пустой список
	if root == nil {
		return [][]int{}
	}

	// Инициализируем очередь и добавляем корень в нее
	queue := []*TreeNode{root}
	result := [][]int{}

	// Пока очередь не пуста
	for len(queue) > 0 {
		// Получаем размер текущей очереди (количество узлов на текущем уровне)
		levelSize := len(queue)
		currentLevel := []int{}

		// Для каждого узла на текущем уровне
		for i := 0; i < levelSize; i++ {
			// Удаляем узел из очереди и добавляем его значение в список текущего уровня
			node := queue[0]
			queue = queue[1:]
			currentLevel = append(currentLevel, node.Val)

			// Если у узла есть левый потомок, добавляем его в очередь
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// Если у узла есть правый потомок, добавляем его в очередь
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Добавляем список текущего уровня в результирующий список
		result = append(result, currentLevel)
	}

	// Возвращаем результирующий список
	return result
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

	result := levelOrder(root)
	fmt.Println(result) // Должен вывести: [[3] [9 20] [15 7]]
}
