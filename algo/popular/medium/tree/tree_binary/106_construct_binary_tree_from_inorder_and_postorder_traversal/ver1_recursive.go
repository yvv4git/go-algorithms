package main

import "fmt"

// TreeNode - структура для представления узла бинарного дерева
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Функция для построения бинарного дерева из inorder и postorder обходов
func buildTree(inorder []int, postorder []int) *TreeNode {
	/*
		METHOD: Recursion

		В postorder обходе (слева->справа->корень) последний элемент всегда является корнем дерева.
		Это обусловлено тем, как мы строим дерево из обходов.

		При построении дерева из inorder и postorder обходов, мы начинаем с последнего элемента в postorder,
		так как он всегда является корнем текущего поддерева.
		Этот элемент разделяет inorder обход на две части: левое поддерево (все элементы слева от корня)
		и правое поддерево (все элементы справа от корня).

		Таким образом, последний элемент в postorder всегда соответствует корню текущего поддерева,
		которое мы строим рекурсивно.

		Inorder:   [9, 3, 15, 20, 7] (слева->корень->справа)
		Postorder: [9, 15, 7, 20, 3] (слева->справа->корень)

		В этом примере последний элемент в postorder - это 3, который является корнем всего дерева.
		В inorder обходе все элементы слева от 3 - это левое поддерево, а все элементы справа от 3 - это правое поддерево.

		TIME COMPLEXITY: O(n^2) в худшем случае, где n - количество узлов в дереве.
		Это происходит потому, что мы ищем индекс корня в inorder обходе, который занимает O(n) времени,
		и для каждого узла мы рекурсивно строим левое и правое поддеревья, что также занимает O(n) времени.

		SPACE COMPLEXITY: O(n) для хранения рекурсивных вызовов, так как в худшем случае мы можем иметь n вызовов рекурсии,
		если дерево сбалансировано.
	*/

	// Если inorder или postorder пусты, дерево пустое
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	// Последний элемент в postorder - корень дерева
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// Ищем индекс корня в inorder
	var indexRoot int
	for i, val := range inorder {
		if val == rootVal {
			indexRoot = i
			break
		}
	}

	// Рекурсивно строим левое и правое поддеревья
	root.Left = buildTree(inorder[:indexRoot], postorder[:indexRoot])
	root.Right = buildTree(inorder[indexRoot+1:], postorder[indexRoot:len(postorder)-1])

	return root
}

// Функция для печати дерева в inorder обходе (слева->корень->справа)
func printInorder(node *TreeNode) {
	if node == nil {
		return
	}
	printInorder(node.Left)
	fmt.Printf("%d ", node.Val)
	printInorder(node.Right)
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)

	// Печатаем дерево в inorder обходе для проверки
	printInorder(root)
}
