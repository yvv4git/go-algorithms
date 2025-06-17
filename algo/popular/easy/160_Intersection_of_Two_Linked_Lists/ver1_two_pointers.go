package main

import "fmt"

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	/*
		APPROACH: Two Pointers approach (intersection detection)
		- Используются два указателя, начинающиеся с голов каждого списка
		- Каждый указатель проходит свой список, затем переключается на другой
		- Если списки пересекаются, указатели встретятся в точке пересечения
		- Если нет пересечения, оба указателя достигнут конца (nil) одновременно

		TIME COMPLEXITY:
		- O(m + n), где m и n - длины списков
		- Каждый указатель проходит оба списка (m + n шагов)
		- В худшем случае оба указателя делают m + n шагов

		SPACE COMPLEXITY:
		- O(1), используется только два дополнительных указателя
		- Не требуется дополнительная память, пропорциональная размеру списков
	*/
	if headA == nil || headB == nil {
		return nil
	}

	// Initialize two pointers
	a, b := headA, headB

	// Loop until we find the intersection or both reach end (nil)
	for a != b {
		// If a reaches end, switch to headB
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		// If b reaches end, switch to headA
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	// Either intersection node or nil
	return a
}

// Explanation:
// The two-pointer approach works by:
// 1. Having each pointer traverse its own list first
// 2. When reaching the end, switching to the other list
// 3. The pointers will meet at the intersection node (if exists)
// because they traverse the same number of nodes (m + n)
// Time complexity: O(m + n)
// Space complexity: O(1)

func main() {
	// Create example linked lists that intersect at node with value 8
	intersectNode := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}

	// List A: 4 -> 1 -> 8 -> 4 -> 5
	listA := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: intersectNode}}

	// List B: 5 -> 6 -> 1 -> 8 -> 4 -> 5
	listB := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: intersectNode}}}

	// Find intersection node
	result := getIntersectionNode(listA, listB)

	if result != nil {
		fmt.Printf("The lists intersect at node with value: %d\n", result.Val)
	} else {
		fmt.Println("The lists do not intersect")
	}
}
