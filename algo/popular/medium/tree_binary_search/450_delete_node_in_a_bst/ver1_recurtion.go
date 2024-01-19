package main

func deleteNode(root *TreeNode, key int) *TreeNode {
	/*
		Method: Recursion
		Time complexity: O(h) / O(log n)
		Space complexity: O(h) / O(log n)

		Time complexity
		Временная сложность будет O(h), где h - высота дерева, так как в худшем случае нам придется пройти по всему дереву.

		Space complexity
		Пространственная сложность будет O(h) для рекурсивного стека вызовов.

		Для решения задачи 450. Delete Node in a BST (удаление узла в двоичном дереве поиска) в языке Go,
		мы будем использовать рекурсивный подход. Временная сложность будет O(h), где h - высота дерева,
		так как в худшем случае нам придется пройти по всему дереву.
		Пространственная сложность будет O(h) для рекурсивного стека вызовов.

		Данный метод называется "рекурсивным удалением узла в двоичном дереве поиска" (Recursive Deletion in Binary Search Tree).
		Метод использует рекурсию для поиска узла, который нужно удалить,
		и затем выполняет необходимые операции для поддержания свойств двоичного дерева поиска после удаления.
	*/
	if root == nil {
		return nil
	}

	// Если ключ меньше корня, идем влево
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		// Если ключ больше корня, идем вправо
		root.Right = deleteNode(root.Right, key)
	} else {
		// Если ключ равен корню, удаляем узел
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Если у узла два потомка, находим минимальный узел в правом поддереве
		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func main() {
	// Пример использования функции deleteNode
	// ...
}
