package _2_reverse_linked_list_II

func reverseBetweenV4(head *ListNode, left int, right int) *ListNode {
	/*
		Method: Two pointers.
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	// Move prev to the node before the node at position left
	for i := 1; i < left; i++ {
		prev = prev.Next
	}

	current := prev.Next

	// Reverse the nodes from left to right
	for i := 0; i < right-left; i++ {
		nextNode := current.Next
		current.Next = nextNode.Next
		nextNode.Next = prev.Next
		prev.Next = nextNode
	}

	return dummy.Next
}
