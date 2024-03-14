package main

import "fmt"

// levelOrderBottom выполняет обход дерева в ширину снизу вверх.
func levelOrderBottom(root *TreeNode) [][]int {
	/*
		METHOD: BFS

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, потому что мы обходим каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае мы можем хранить в очереди все узлы одного уровня дерева.
	*/
	// Если корень дерева пуст, возвращаем пустой срез.
	if root == nil {
		return [][]int{}
	}

	// result будет содержать результат обхода уровней дерева.
	var result [][]int
	// queue - это очередь для обхода узлов дерева.
	queue := []*TreeNode{root}

	// Пока очередь не пуста, выполняем обход уровней.
	for len(queue) > 0 {
		// levelSize - количество узлов на текущем уровне.
		levelSize := len(queue)
		// currentLevel - список значений узлов на текущем уровне.
		var currentLevel []int

		// Проходим по каждому узлу на текущем уровне.
		for i := 0; i < levelSize; i++ {
			// Извлекаем первый узел из очереди.
			node := queue[0]
			// Удаляем первый узел из очереди.
			queue = queue[1:]
			// Добавляем значение узла в список значений текущего уровня.
			currentLevel = append(currentLevel, node.Val)

			// Если у текущего узла есть левый потомок, добавляем его в очередь.
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			// Если у текущего узла есть правый потомок, добавляем его в очередь.
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Добавляем список значений текущего уровня в начало результата.
		result = append([][]int{currentLevel}, result...)
	}

	// Возвращаем результат - список значений узлов дерева, обходом снизу вверх.
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

	levels := levelOrderBottom(root)
	for _, level := range levels {
		fmt.Println(level)
	}
}
