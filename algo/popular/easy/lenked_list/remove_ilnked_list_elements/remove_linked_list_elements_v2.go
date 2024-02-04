package remove_ilnked_list_elements

func removeElementsV2(head *ListNode, val int) *ListNode {
	/*
		METHOD: Loop
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	if head == nil {
		return nil
	}

	for head != nil && head.Val == val {
		head = head.Next
	}

	current, prev := head, head

	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}

		current = current.Next
	}

	return head
}
