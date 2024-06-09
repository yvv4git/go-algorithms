package _816_double_number_represented_as_linked_list

func doubleItV2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Val >= 5 {
		newNode := &ListNode{Val: 1, Next: head}
		head = newNode
	}

	curr := head
	for curr != nil {
		if curr.Val < 5 {
			curr.Val = (curr.Val * 2) % 10
		} else {
			curr.Val = (curr.Val*2 + 1) % 10
		}
		curr = curr.Next
	}

	return head
}
