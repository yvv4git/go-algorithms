package main

import "math"

// Функция для обхода дерева в глубину без рекурсии
func getMinimumDifferenceV3(root *TreeNode) int {
	/*
		METHOD: DFS iterative
		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, поскольку мы посещаем каждый узел ровно один раз.
		Space complexity: O(n), так как в худшем случае мы можем поместить все узлы в стек.
	*/
	// Создаем пустой стек для хранения узлов
	stack := make([]*TreeNode, 0)
	// Устанавливаем текущий узел в корень дерева
	cur := root
	// Инициализируем предыдущий узел как nil
	var pre *TreeNode
	// Инициализируем минимальную разницу как максимально возможное значение
	min := math.MaxInt32

	// Пока стек не пуст или есть узел для обработки
	for len(stack) != 0 || cur != nil {
		// Пока есть узел для обработки
		for cur != nil {
			// Добавляем узел в стек
			stack = append(stack, cur)
			// Переходим к левому поддереву
			cur = cur.Left
		}
		// Если предыдущий узел не nil и разница между текущим узлом и предыдущим меньше минимальной разницы
		if pre != nil && min > stack[len(stack)-1].Val-pre.Val {
			// Обновляем минимальную разницу
			min = stack[len(stack)-1].Val - pre.Val
		}
		// Обновляем предыдущий узел
		pre = stack[len(stack)-1]
		// Переходим к правому поддереву
		cur = stack[len(stack)-1].Right
		// Удаляем последний узел из стека
		stack = stack[:len(stack)-1]
	}

	// Возвращаем минимальную разницу
	return min
}
