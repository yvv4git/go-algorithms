package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

var head *ListNode

func countNodes(head *ListNode) int {
	count := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		count++
	}

	return count
}

func sortedListToBSTRec(n int) *TreeNode {
	if n <= 0 {
		return nil
	}

	left := sortedListToBSTRec(n / 2)

	root := &TreeNode{Val: head.Val}
	root.Left = left

	head = head.Next

	root.Right = sortedListToBSTRec(n - n/2 - 1)

	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	/*
		METHOD: Morris Traversal
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: O(1)

		Time complexity
		Временная сложность метода "Преобразование в BST с помощью Morris Traversal" составляет O(n log n) в худшем случае,
		когда дерево является сбалансированным. Это связано с тем, что для каждого узла мы выполняем поиск среднего узла,
		что требует O(n) времени.

		Space complexity
		Пространственная сложность метода "Преобразование в BST с помощью Morris Traversal" составляет O(1) в худшем случае,
		так как мы не используем дополнительное пространство, кроме необходимого для узлов BST.
	*/
	n := countNodes(head)
	return sortedListToBSTRec(n)
}

func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

func push(head *ListNode, data int) {
	newNode := &ListNode{Val: data, Next: head}
	head = newNode
}

func main() {
	head = &ListNode{Val: 7}
	push(head, 6)
	push(head, 5)
	push(head, 4)
	push(head, 3)
	push(head, 2)
	push(head, 1)

	root := sortedListToBST(head)
	preOrder(root)
}
