package main

// Функция для перестановки узлов в списке так, чтобы он стал палиндромом
func reorderListV1(head *ListNode) {
	/*
		METHOD: Fast & slow pointers (Двухпутевой указатель)
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	if head == nil || head.Next == nil {
		return
	}

	// Находим середину списка
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Разворачиваем вторую половину списка
	prev.Next = nil
	secondHalf := reverseList(slow)

	// Переставляем узлы в списке так, чтобы он стал палиндромом
	for head != nil && secondHalf != nil {
		next := head.Next
		head.Next = secondHalf
		secondHalf = secondHalf.Next
		head.Next.Next = next
		head = next
	}
}

// Функция для разворота списка
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
