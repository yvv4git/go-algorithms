//go:build ignore

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	/*
		APPROACH: Hash set (map)
		- Используется hash set (map), чтобы хранить уже посещённые узлы.
		- При проходе по списку для каждого узла проверяем, был ли он уже посещён.
		- Если узел уже есть в hash set — значит, найден цикл.
		- Если дошли до конца списка (nil), значит цикла нет.

		TIME COMPLEXITY:
		- O(n), где n — количество элементов в списке. Каждый узел посещается не более одного раза.

		SPACE COMPLEXITY:
		- O(n), так как в худшем случае в hash set попадут все узлы списка (если цикла нет).
	*/
	visited := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := visited[head]; ok {
			return true
		}
		visited[head] = struct{}{}
		head = head.Next
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
