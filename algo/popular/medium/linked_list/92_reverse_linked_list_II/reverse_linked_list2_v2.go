package _2_reverse_linked_list_II

func reverseBetweenV2(head *ListNode, left int, right int) *ListNode {
	/*
		METHOD: Recursion.
		TIME COMPLEXITY: O(n).
		SPACE COMPLEXITY: O(1).
	*/

	// Из условия задачи, надо поменять узлы местами, когда left <= right.
	if head == nil || left >= right {
		return head
	}

	var reverseRecursion func(node *ListNode, m int) *ListNode

	reverseRecursion = func(node *ListNode, m int) *ListNode {
		if m == left { // Определяем, достигли ли мы узла, который надо reverse.
			var prev, current *ListNode = nil, node
			for m <= right { // Как только дошли до левого элемента, то меняем ноды текущую с предыдущей.
				current.Next, prev, current = prev, current, current.Next
				m++
			}
			node.Next = current

			return prev
		} else if m < left { // Если мы не достигли узла, который надо reverse, то переходим к следующему узлу и увеличиваем счетчик m.
			node.Next = reverseRecursion(node.Next, m+1)
		}

		return node
	}

	// Начинаем рекурсивно двигаться по списку(Linked List).
	return reverseRecursion(head, 1)
}
