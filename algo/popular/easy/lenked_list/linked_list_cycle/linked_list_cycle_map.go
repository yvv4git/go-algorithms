package main

func hasCycleWithMap(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}

	return false
}

// ListNode - ...
type ListNode struct {
	Val  int
	Next *ListNode
}
