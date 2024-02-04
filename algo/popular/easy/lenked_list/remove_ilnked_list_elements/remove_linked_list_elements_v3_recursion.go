package remove_ilnked_list_elements

func removeElementsV3(head *ListNode, val int) *ListNode {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)
	*/
	if head == nil {
		return nil
	}

	head.Next = removeElementsV3(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}
