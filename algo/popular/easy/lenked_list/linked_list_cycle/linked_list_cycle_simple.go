package main

func hasCycleSimple(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	p1, p2 := head, head.Next
	for p1 != p2 {
		if p2 == nil || p2.Next == nil {
			return false
		}

		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return true
}

/*// ListNode - used as node
type ListNode struct {
	Val  int
	Next *ListNode
}
*/
