//go:build ignore

package main

import "fmt"

/**
 * Определение узла бинарного дерева.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTargetCopy(original *TreeNode, cloned *TreeNode, target *TreeNode) *TreeNode {
	// Стек для обхода оригинального и клонированного дерева
	stackOriginal := []*TreeNode{original}
	stackCloned := []*TreeNode{cloned}

	for len(stackOriginal) > 0 {
		// Извлекаем текущие узлы из стеков
		nodeOriginal := stackOriginal[len(stackOriginal)-1]
		stackOriginal = stackOriginal[:len(stackOriginal)-1]

		nodeCloned := stackCloned[len(stackCloned)-1]
		stackCloned = stackCloned[:len(stackCloned)-1]

		// Если текущий узел оригинального дерева равен целевому узлу,
		// возвращаем соответствующий узел в клонированном дереве
		if nodeOriginal == target {
			return nodeCloned
		}

		// Добавляем правые и левые поддеревья в стек
		if nodeOriginal.Right != nil {
			stackOriginal = append(stackOriginal, nodeOriginal.Right)
			stackCloned = append(stackCloned, nodeCloned.Right)
		}
		if nodeOriginal.Left != nil {
			stackOriginal = append(stackOriginal, nodeOriginal.Left)
			stackCloned = append(stackCloned, nodeCloned.Left)
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
