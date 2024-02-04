package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(h), где h - высота дерева.
		Space complexity: O(h)

		Time complexity
		Временная сложность функции deleteNodeV2() в худшем случае составляет O(h), где h - высота дерева.
		Это связано с тем, что в худшем случае функция может потребоваться для обхода всего дерева, если удаляемый узел является листом.

		Space complexity
		Пространственная сложность функции также составляет O(h), так как в худшем случае она может хранить в стеке вызовов h рекурсивных вызовов.

		В данном решении используется подход на основе рекурсии, который является стандартным для решения задач, связанных с бинарными деревьями поиска.
		Если поддерево не пустое, функция сравнивает значение нового узла с значением корневого узла.
		В зависимости от результата сравнения функция рекурсивно вызывает себя для левого или правого поддерева,
		передавая в качестве корня соответствующее поддерево.
		Таким образом, функция находит место для вставки нового узла, создает его и вставляет в дерево, сохраняя свойство бинарного дерева поиска.
	*/

	// Если дерево пустое, создаем новый узел и возвращаем его
	if root == nil {
		return &TreeNode{Val: val}
	}

	// Если значение нового узла меньше корневого, добавляем его в левое поддерево
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		// Иначе добавляем его в правое поддерево
		root.Right = insertIntoBST(root.Right, val)
	}

	// Возвращаем корень дерева
	return root
}
