package _816_double_number_represented_as_linked_list

func doubleItV1(head *ListNode) *ListNode {
	/*
		Method:
		Time complexity: O(n)
		Space complexity: O(1)

		Основная идея метода заключается в том, чтобы удвоить значение каждого узла в списке, если его значение меньше 5,
		и удвоить значение и прибавить 1, если значение узла больше или равно 5.
		Если значение первого узла больше или равно 5, то создается новый узел со значением 1 и он становится первым узлом списка.
	*/
	result := head

	if head.Val >= 5 {
		head = &ListNode{Val: 1, Next: head}
		result = head
		head = head.Next
	}

	for head != nil && head.Next != nil {

		if head.Next.Val < 5 {
			head.Val = (head.Val * 2) % 10
		} else {
			head.Val = (head.Val*2 + 1) % 10
		}

		head = head.Next
	}

	head.Val = (head.Val * 2) % 10

	return result
}
