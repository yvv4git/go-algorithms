package main

import "math"

// Функция isValidBSTV4 проверяет, является ли данное бинарное дерево деревом поиска.
// Дерево поиска - это бинарное дерево, в котором для каждого узла выполняется условие:
// все узлы в левом поддереве меньше значения узла, а все узлы в правом поддереве больше значения узла.
func isValidBSTV4(root *TreeNode) bool {
	/*
		METHOD: DFS
		Time complexity: O(n)
		Space complexity: O(1)
	*/

	// Определяем вложенную функцию validate, которая будет проверять каждый узел дерева.
	// Она принимает три аргумента: узел, минимальное и максимальное значения.
	var validate func(n *TreeNode, min int, max int) bool

	validate = func(n *TreeNode, min int, max int) bool {
		// Если узел пустой, то он является валидным.
		if n == nil {
			return true
		}
		// Если значение узла меньше минимального или больше максимального, то дерево не является валидным.
		if min >= n.Val || max <= n.Val {
			return false
		}
		// Проверяем левое и правое поддеревья.
		// Для левого поддерева мы устанавливаем новый максимальный предел - значение текущего узла.
		// Для правого поддерева мы устанавливаем новый минимальный предел - значение текущего узла.
		return validate(n.Left, min, n.Val) && validate(n.Right, n.Val, max)
	}

	// Начинаем проверку с корня дерева, минимальное значение - минимально возможное, максимальное - максимально возможное.
	return validate(root, math.MinInt, math.MaxInt)
}
