package main

import (
	"fmt"
	"math"
)

// Функция для нахождения максимальной суммы пути в дереве
func maxPathSum(root *TreeNode) int {
	/*
		METHOD: DFS post-order traversal
		В этом коде мы используем подход, известный как "пост-обход" (post-order traversal),
		чтобы рекурсивно обойти дерево, начиная с листьев и двигаясь к корню.
		Мы используем вспомогательную функцию maxGain, которая возвращает максимальную сумму пути,
		которая может быть получена из узла и одного из его поддеревьев.
		Мы обновляем глобальную переменную maxSum при обнаружении более выгодного пути.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, потому что мы посещаем каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n), так как в худшем случае (когда дерево сбалансировано) глубина рекурсии может достигать n.
		В худшем случае (когда дерево несбалансировано) мы можем хранить n рекурсивных вызовов в стеке.
	*/
	maxSum := math.MinInt32 // Инициализируем максимальную сумму как минимальное значение int32

	// Вспомогательная функция для поиска максимальной суммы пути
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0 // Если узел пустой, возвращаем 0
		}

		// Рекурсивно находим максимальную сумму в левом и правом поддереве
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// Обновляем максимальную сумму, если текущая сумма больше
		priceNewPath := node.Val + leftGain + rightGain
		maxSum = max(maxSum, priceNewPath)

		// Возвращаем максимальную сумму для одного из поддеревьев, добавляя значение текущего узла
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root) // Запускаем функцию maxGain с корнем дерева

	return maxSum // Возвращаем максимальную сумму пути
}

// Вспомогательная функция для нахождения максимума из двух чисел
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// Создаем дерево для примера
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	// Находим и выводим максимальную сумму пути
	fmt.Println(maxPathSum(root)) // Вывод: 6
}
