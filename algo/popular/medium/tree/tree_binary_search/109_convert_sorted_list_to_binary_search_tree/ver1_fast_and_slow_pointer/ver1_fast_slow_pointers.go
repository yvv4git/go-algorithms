package main

import "fmt"

// Функция для поиска среднего элемента в списке
// Входные данные: указатель на начало списка
// Выходные данные: указатель на средний элемент списка
// Временная сложность: O(n), где n - количество элементов в списке
// Пространственная сложность: O(1), так как мы используем постоянное количество переменных
func findMiddle(head *ListNode) *ListNode {
	var prev *ListNode
	slow, fast := head, head
	// Используем алгоритм "Floyd's Cycle Detection Algorithm" для поиска среднего элемента
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// Отделяем список на две части
	if prev != nil {
		prev.Next = nil
	}
	return slow
}

// Функция для преобразования отсортированного списка в бинарное дерево поиска
// Входные данные: указатель на начало списка
// Выходные данные: указатель на корень бинарного дерева поиска
// Временная сложность: O(n log n), так как на каждом уровне рекурсии мы делим список на пополам, а затем рекурсивно обрабатываем каждую половину
// Пространственная сложность: O(log n), так как мы используем дополнительное пространство для стека вызовов рекурсии
func sortedListToBST(head *ListNode) *TreeNode {
	/*
		METHOD: Fast and slow pointers
		TIME COMPLEXITY: O(n log n), так как на каждом уровне рекурсии мы делим список на пополам, а затем рекурсивно обрабатываем каждую половину
		Space complexity: O(log n), так как мы используем дополнительное пространство для стека вызовов рекурсии
	*/
	// Если список пуст, то и дерево пусто
	if head == nil {
		return nil
	}
	// Находим средний элемент списка
	mid := findMiddle(head)
	// Создаем новую вершину дерева с значением среднего элемента
	node := &TreeNode{Val: mid.Val}
	// Если список состоит из одного элемента, то дерево состоит из одной вершины
	if head == mid {
		return node
	}
	// Рекурсивно преобразуем левое и правое поддеревья
	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(mid.Next)
	return node
}

// Функция для вывода дерева
func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9, Next: nil}}}}}
	root := sortedListToBST(head)
	printTree(root)
}
