package main

func partitionRecursive(head *ListNode, x int) *ListNode {
	/*
		METHOD: Recursive

		TIME COMPLEXITY: O(n), т.к. в худшем случае, когда все узлы списка больше или равны x,
		глубина рекурсии будет равна количеству узлов в списке, что дает временную сложность O(n).

		SPACE COMPLEXITY: O(n) в худшем случае, когда глубина рекурсии равна количеству узлов в списке, так как для каждого узла в стеке сохраняется его состояние.
	*/
	if head == nil || head.Next == nil {
		return head
	}

	head.Next = partitionRecursive(head.Next, x)

	if head.Val < x {
		return head
	} else {
		next := head.Next
		head.Next = nil
		return next
	}
}
