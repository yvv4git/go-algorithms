package remove_ilnked_list_elements

func removeElementsV1(head *ListNode, val int) *ListNode {
	/*
		Method: Loop
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	var (
		rear = head          // pointer to the next node of the current node
		pre  = new(ListNode) // dummy head node
		cur  = pre           // pointer to the current node (the node before the rear node)
	)
	pre.Next = head // link the dummy head node to the head node

	for rear != nil { // loop until the rear node is nil
		if rear.Val == val { // if the rear node's value is equal to the given value
			cur.Next = cur.Next.Next // delete the rear node
			rear = cur.Next          // move the rear node to the next node of the current node
		} else {
			cur = rear       // move the current node to the rear node
			rear = rear.Next // move the rear node to the next node of the rear node
		}
	}

	// return the next node of the dummy head node
	return pre.Next
}
