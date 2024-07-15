package main

import "fmt"

// Определение структуры TreeNode для представления узла бинарного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция createBinaryTreeFromDescriptions создает бинарное дерево из списка описаний
func createBinaryTreeFromDescriptions(descriptions [][]int) *TreeNode {
	/*
		METHOD: Map/Dictionary with Relationship graph

		TIME COMPLEXITY: O(n), где n — количество описаний. Мы проходим по всем описаниям один раз.

		SPACE COMPLEXITY: O(n), где n — количество уникальных узлов в дереве. Мы храним все узлы в nodeMap и дополнительно в parentMap.
	*/
	// Создаем карту для хранения узлов по их значениям
	nodeMap := make(map[int]*TreeNode)
	// Создаем карту для хранения родительских узлов
	parentMap := make(map[int]bool)

	// Проходим по всем описаниям
	for _, desc := range descriptions {
		parentVal, childVal, isLeft := desc[0], desc[1], desc[2]

		// Получаем или создаем узлы для родителя и ребенка
		parentNode, exists := nodeMap[parentVal]
		if !exists {
			parentNode = &TreeNode{Val: parentVal}
			nodeMap[parentVal] = parentNode
		}

		childNode, exists := nodeMap[childVal]
		if !exists {
			childNode = &TreeNode{Val: childVal}
			nodeMap[childVal] = childNode
		}

		// Устанавливаем связь между родителем и ребенком
		if isLeft == 1 {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}

		// Отмечаем ребенка как имеющего родителя
		parentMap[childVal] = true
	}

	// Находим корневой узел, который не имеет родителя
	var root *TreeNode
	for val, node := range nodeMap {
		if !parentMap[val] {
			root = node
			break
		}
	}

	return root
}

func main() {
	descriptions := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
	root := createBinaryTreeFromDescriptions(descriptions)
	fmt.Println(root.Val) // Вывод: 50
}
