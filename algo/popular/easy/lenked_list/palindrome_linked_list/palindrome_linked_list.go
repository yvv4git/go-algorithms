package main

func isPalindrome(head *ListNode) bool {
	var values []int

	cur := head
	for cur != nil {
		values = append(values, cur.Val)
		cur = cur.Next
	}

	i := 0
	j := len(values) - 1
	for i < j {
		if values[i] != values[j] {
			return false
		}

		i++
		j--
	}

	return true
}

// ListNode - implementation of Node
type ListNode struct {
	Val  int
	Next *ListNode
}
