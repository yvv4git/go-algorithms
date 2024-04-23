package main

import "fmt"

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	/*
		METHOD: Recursion

		TIME COMPLEXITY: O(n), где n - общее количество узлов в обоих деревьях, потому что мы обходим каждый узел ровно один раз.

		SPACE COMPLEXITY: O(n) в худшем случае, когда результирующее дерево будет иметь то же количество узлов, что и самое глубокое из исходных деревьев.
		Это происходит, когда одно из деревьев полностью находится внутри другого.
		В идеальном случае, когда деревья сбалансированы, пространственная сложность будет O(log n), где log n - это глубина самого глубокого дерева.
	*/
	// Если оба корня пусты, результат объединения также будет пустым
	if root1 == nil && root2 == nil {
		return nil
	}

	// Создаем новый узел для результирующего дерева
	newRoot := &TreeNode{}

	// Если только один из корней не пустой, прибавляем значение этого корня к новому узлу
	if root1 != nil {
		newRoot.Val += root1.Val
	}
	if root2 != nil {
		newRoot.Val += root2.Val
	}

	// Рекурсивно объединяем левое и правое поддерево
	// Если одно из деревьев не имеет левого/правого потомка, передаем nil
	newRoot.Left = mergeTrees(getChildNode(root1, true), getChildNode(root2, true))
	newRoot.Right = mergeTrees(getChildNode(root1, false), getChildNode(root2, false))

	return newRoot
}

// Вспомогательная функция для безопасного получения потомка узла
func getChildNode(root *TreeNode, isLeft bool) *TreeNode {
	if root != nil {
		if isLeft {
			return root.Left
		}
		return root.Right
	}
	return nil
}

func main() {
	// Пример использования функции mergeTrees
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 2}}
	t2 := &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 7}}}

	mergedTree := mergeTrees(t1, t2)
	fmt.Println(mergedTree) // Вывод: указатель на корень объединенного дерева
}
