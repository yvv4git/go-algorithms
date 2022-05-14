package main

func reverseList(head *ListNode) *ListNode {
	// Если у нас только одна нода, то нет смысла продолжать дальше
	if head == nil || head.Next == nil {
		return head
	}

	// Начиная с 2-х нод надо перестраивать
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev
}

// ListNode - ...
type ListNode struct {
	Val  int
	Next *ListNode
}
