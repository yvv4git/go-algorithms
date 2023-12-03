package main

import (
	"fmt"
	"testing"
)

/*
Визуализация того, как работает метод Fast & slow pointers.
*/

func TestFastAndSlowPointers01(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}

	// 1 -> 2 -> 3 -> 4 -> 5
	// Очевидно, что 3 здесь середина.
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	middle := middleNode(node1)
	fmt.Println(middle.Val) // 3
}

func TestFastAndSlowPointers02(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}

	// 1 -> 2 -> 3 -> 4 -> 5 -> 6
	// Очевидно, что 4 здесь середина.
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6

	middle := middleNode(node1)
	fmt.Println(middle.Val) // 4
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
