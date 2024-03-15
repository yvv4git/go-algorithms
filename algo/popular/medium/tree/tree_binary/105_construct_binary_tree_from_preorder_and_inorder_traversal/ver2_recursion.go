package main

// Функция для построения бинарного дерева из предварительного и симметричного обходов
func buildTreeRecursion(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// Первый элемент предварительного обхода - корень дерева
	root := &TreeNode{Val: preorder[0]}

	// Находим индекс корня в симметричном обходе, чтобы разделить левое и правое поддеревья
	var rootIndex int
	for i, val := range inorder {
		if val == root.Val {
			rootIndex = i
			break
		}
	}

	// Рекурсивно строим левое и правое поддеревья
	root.Left = buildTreeRecursion(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTreeRecursion(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

//// Функция для печати дерева в симметричном порядке
//func printInorder(node *TreeNode) {
//	if node == nil {
//		return
//	}
//	printInorder(node.Left)
//	fmt.Printf("%d ", node.Val)
//	printInorder(node.Right)
//}

//func main() {
//	preorder := []int{3, 9, 20, 15, 7}
//	inorder := []int{9, 3, 15, 20, 7}
//
//	root := buildTree(preorder, inorder)
//
//	// Печать дерева в симметричном порядке для проверки
//	printInorder(root)
//}
