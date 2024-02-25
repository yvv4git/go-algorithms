package main

// Функция levelOrderV2 реализует обход дерева в порядке уровней с использованием рекурсии.
func levelOrderV2(root *TreeNode) [][]int {
	/*
		METHOD: DFS with level recording

		TIME COMPLEXITY: O(n), где n - количество узлов в дереве. Это связано с тем, что мы посещаем каждый узел ровно один раз.
		Однако в худшем случае, когда дерево является сбалансированным,
		глубина рекурсии будет равна высоте дерева (O(log N) для сбалансированного дерева),
		что приводит к дополнительной пространственной сложности O(log N) на каждый уровень рекурсии.

		SPACE COMPLEXITY: O(n), так как в очереди находится по одному узлу на каждом уровне, включая последний.
		В среднем, когда дерево не является сбалансированным, пространственная сложность может быть меньше,
		но в любом случае будет меньше линейной, так как мы не храним узлы, которые уже были посещены.
	*/
	// Инициализация слайса для хранения результата.
	var ans [][]int
	// Если корень дерева равен nil, возвращаем пустой слайс.
	if root == nil {
		return ans
	}
	// Вызов вспомогательной функции helper для обхода дерева.
	helper(root, 0, &ans)
	// Возвращаем результат.
	return ans
}

// Вспомогательная функция helper рекурсивно обходит дерево и заполняет результирующий слайс ans.
func helper(node *TreeNode, level int, ans *[][]int) {
	// Если текущий узел равен nil, выходим из рекурсии.
	if node == nil {
		return
	}
	// Если текущий уровень равен длине слайса ans, добавляем новый слайс для хранения значений узлов этого уровня.
	if len(*ans) == level {
		*ans = append(*ans, []int{})
	}
	// Добавляем значение текущего узла в соответствующий слайс на текущем уровне.
	(*ans)[level] = append((*ans)[level], node.Val)
	// Рекурсивный вызов для левого и правого потомков текущего узла.
	helper(node.Left, level+1, ans)
	helper(node.Right, level+1, ans)
}
