package main

import "fmt"

// Функция для нахождения диаметра бинарного дерева
func diameterOfBinaryTree(root *TreeNode) int {
	/*
		METHOD: Post-order traversal
		Этот подход позволяет нам вычислить высоту каждого поддерева и сразу же найти диаметр.
		В данном случае, мы используем пост-обход для вычисления высоты каждого поддерева.
		Когда мы посещаем узел, мы сначала вычисляем высоту его левого и правого поддеревьев.
		Затем мы обновляем максимальный диаметр, который является суммой высот левого и правого поддеревьев.

		Далее, мы возвращаем высоту текущего поддерева, которая равна максимуму высоты левого и правого поддеревьев, плюс один (текущий узел).
		Это позволяет нам вычислить высоту каждого поддерева, поскольку мы рекурсивно вычисляем высоту каждого поддерева и обновляем максимальный диаметр.

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве, так как мы посещаем каждый узел ровно один раз.

		SPACE COMPLEXITY: O(h), где h - высота дерева, так как нам необходимо хранить в стеке рекурсивные вызовы до тех пор,
		пока не достигнем листа.
	*/
	maxDiameter := 0 // Переменная для хранения максимального диаметра

	// Вспомогательная функция для вычисления высоты поддерева
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// Рекурсивно вычисляем высоту левого и правого поддеревьев
		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)

		// Обновляем максимальный диаметр
		maxDiameter = max(maxDiameter, leftDepth+rightDepth)

		// Возвращаем максимальную высоту между левой и правой глубиной, плюс один (текущий узел)
		return max(leftDepth, rightDepth) + 1
	}

	// Вызываем вспомогательную функцию для корня дерева
	depth(root)

	// Возвращаем максимальный диаметр
	return maxDiameter
}

// Вспомогательная функция для нахождения максимума
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример использования
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(diameterOfBinaryTree(root)) // Вывод: 3
}
