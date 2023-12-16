package _1_rotate_list

func rotateRightV2(head *ListNode, k int) *ListNode {
	/*
		Method: Length Calculation Method
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность этого решения O(n), где n - количество узлов в списке,
		так как мы проходим по списку дважды: один раз для вычисления длины и один раз для поиска нового начала списка.

		Space complexity
		Пространственная сложность O(1), так как мы используем константное количество дополнительной памяти для хранения указателей.

		В этом решении мы сначала вычисляем длину списка.
		Затем мы вычисляем фактическое количество поворотов, используя остаток от деления k на длину списка.
		Если фактическое количество поворотов равно 0, мы возвращаем исходный список, так как он не нуждается в повороте.
		В противном случае мы находим новый начало списка, используя цикл.
		Затем мы поворачиваем список, изменяя указатели следующих узлов.
	*/
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Вычисляем длину списка
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}

	// Вычисляем фактическое количество поворотов
	k = k % length
	if k == 0 {
		return head
	}

	// Находим новый начало списка
	newHead := head
	for i := 0; i < length-k-1; i++ {
		newHead = newHead.Next
	}

	// Поворачиваем список
	tail := newHead
	newHead = newHead.Next
	tail.Next = nil
	curr = newHead
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = head

	return newHead
}
