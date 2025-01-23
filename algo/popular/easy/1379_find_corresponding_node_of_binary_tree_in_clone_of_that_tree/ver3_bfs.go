package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTargetCopy(original *TreeNode, cloned *TreeNode, target *TreeNode) *TreeNode {
	// Очереди для обхода оригинального и клонированного дерева
	queueOriginal := []*TreeNode{original}
	queueCloned := []*TreeNode{cloned}

	for len(queueOriginal) > 0 {
		// Извлекаем текущие узлы из очередей
		nodeOriginal := queueOriginal[0]
		queueOriginal = queueOriginal[1:]

		nodeCloned := queueCloned[0]
		queueCloned = queueCloned[1:]

		// Если текущий узел оригинального дерева равен целевому узлу,
		// возвращаем соответствующий узел в клонированном дереве
		if nodeOriginal == target {
			return nodeCloned
		}

		// Добавляем левые и правые поддеревья в очередь
		if nodeOriginal.Left != nil {
			queueOriginal = append(queueOriginal, nodeOriginal.Left)
			queueCloned = append(queueCloned, nodeCloned.Left)
		}
		if nodeOriginal.Right != nil {
			queueOriginal = append(queueOriginal, nodeOriginal.Right)
			queueCloned = append(queueCloned, nodeCloned.Right)
		}
	}

	// Если узел не найден, возвращаем nil
	return nil
}

func main() {
	// Создаем оригинальное дерево
	original := &TreeNode{Val: 1}
	original.Left = &TreeNode{Val: 2}
	original.Right = &TreeNode{Val: 3}
	original.Left.Left = &TreeNode{Val: 4}
	original.Left.Right = &TreeNode{Val: 5}

	// Клонируем дерево (в реальности клон будет создан отдельно, здесь для примера используем то же дерево)
	cloned := original

	// Целевой узел в оригинальном дереве (например, узел со значением 5)
	target := original.Left.Right

	// Находим соответствующий узел в клоне
	result := getTargetCopy(original, cloned, target)
	fmt.Println(result.Val) // Вывод: 5
}
