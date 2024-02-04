package main

var first, second, prev *TreeNode

func recoverTreeV1(root *TreeNode) {
	/*
		METHOD: Inorder
		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы проходим по всем узлам.
		SPACE COMPLEXITY: O(h), где h - высота дерева, так как в худшем случае мы можем хранить в стеке все узлы дерева.

		Метод решения задачи:
		1. Инициализируем два указателя на предыдущий и текущий узлы.
		2. Запускаем рекурсивный обход дерева (например, In-Order обход).
		3. В процессе обхода, мы сравниваем текущий узел с предыдущим узлом.
		Если текущий узел меньше предыдущего узла, то мы нашли два узла, которые были помечены неправильно.
		4. Мы также должны отслеживать первый узел, который мы нашли, так как он может быть первым узлом, который нарушает инвариант BST.
		5. После обхода, мы должны поменять значения двух узлов, которые были помечены неправильно.
	*/
	first, second, prev = nil, nil, nil
	inorder(root)
	first.Val, second.Val = second.Val, first.Val
}

// Функция inorder выполняет обход дерева в порядке "inorder" (слева-корень-право).
// Это означает, что она будет посещать узлы в следующем порядке: левое поддерево, узел, правое поддерево.
func inorder(root *TreeNode) {
	/*
		В этом коде, функция inorder использует DFS для обхода дерева.
		Она начинает с корня дерева и спускается вниз по левому поддереву, пока не достигнет листа.
		Затем она обходит узел (обрабатывает его), и затем спускается вниз по правому поддереву.

		В вашем случае, функция inorder используется для обхода дерева в порядке "inorder", но не для поиска двух узлов,
		которые нарушают порядок возрастания. Вместо этого, она используется для обхода дерева и поиска двух узлов,
		которые нарушают порядок возрастания.
	*/
	// Если узел пустой, то мы дошли до листа дерева и возвращаемся назад.
	if root == nil {
		return
	}

	// Рекурсивный вызов функции для левого поддерева.
	inorder(root.Left)

	// Если предыдущий узел не пустой и значение текущего узла меньше значения предыдущего узла,
	// то мы нашли два узла, которые нарушают порядок возрастания.
	if prev != nil && root.Val < prev.Val {
		// Если первый узел еще не был найден, то мы его запоминаем.
		if first == nil {
			first = prev
		}
		// В любом случае мы запоминаем второй узел.
		second = root
	}

	// После того, как мы обработали текущий узел, мы запоминаем его как предыдущий.
	prev = root

	// Рекурсивный вызов функции для правого поддерева.
	inorder(root.Right)
}
