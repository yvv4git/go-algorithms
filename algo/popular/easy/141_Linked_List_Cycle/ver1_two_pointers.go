//go:build ignore

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	/*
		APPROACH: Two Pointers approach (slow & fast)
		- Используются два указателя: медленный (slow) и быстрый (fast).
		- Оба начинаются с головы списка.
		- На каждом шаге slow двигается на один элемент, а fast — на два.
		- Если в списке есть цикл, fast и slow обязательно встретятся внутри цикла.
		- Если fast дойдет до конца списка (null), значит цикла нет.

		TIME COMPLEXITY:
		- O(n), где n — количество элементов в списке. Каждый указатель проходит не более чем по всем элементам.

		SPACE COMPLEXITY:
		- O(1), так как используются только два дополнительных указателя вне зависимости от размера списка.
	*/
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}

func main() {
	// Пример 1: список с циклом [3,2,0,-4], tail -> 2-й элемент
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2 // цикл
	fmt.Println("Список 1 (ожидается true):", hasCycle(node1))

	// Пример 2: список без цикла [1,2]
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n1.Next = n2
	fmt.Println("Список 2 (ожидается false):", hasCycle(n1))
}
