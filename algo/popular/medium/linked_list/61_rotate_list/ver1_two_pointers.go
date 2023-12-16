package _1_rotate_list

// RotateRight поворачивает список вправо на k позиций.
func rotateRightV1(head *ListNode, k int) *ListNode {
	/*
		Method: Two-Pointer
		Time complexity: O(n), где n - количество узлов в списке.
		Space complexity: O(1).
	*/

	// Если список пуст, содержит один узел или k равно 0, то возвращаем исходный список.
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Вычисляем длину списка.
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	// Вычисляем фактическое количество поворотов.
	k = k % length
	if k == 0 {
		return head
	}

	// Находим новый хвост и новую голову повернутого списка.
	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// Поворачиваем список.
	newHead := slow.Next
	slow.Next = nil
	fast.Next = head

	return newHead
}
