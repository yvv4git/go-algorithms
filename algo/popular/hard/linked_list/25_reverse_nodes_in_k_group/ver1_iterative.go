package _5_reverse_nodes_in_k_group

// Функция для переворачивания группы узлов в связном списке
func reverseKGroupV1(head *ListNode, k int) *ListNode {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		Time complexity
		Временная сложность O(n), где n - количество узлов в списке.
		Мы проходим по списку дважды: один раз для определения начала группы, а второй раз для переворачивания группы.

		Space complexity
		Пространственная сложность O(1), так как мы не используем дополнительную память, кроме некоторых переменных для указателей
	*/

	// Создаем фиктивный узел, который будет указывать на начало списка
	dummy := &ListNode{Next: head}
	prevGroup := dummy

	// Проходим по списку, пока есть достаточно узлов для переворачивания
	for {
		kthNode := getKthNode(prevGroup, k)
		if kthNode == nil {
			// Если узлов недостаточно, то выходим из цикла
			break
		}

		// Сохраняем следующий узел группы
		nextGroup := kthNode.Next

		// Переворачиваем группу
		reverseList(prevGroup, nextGroup)

		// Обновляем указатель на начало следующей группы
		prevGroup, kthNode.Next = kthNode, nextGroup
	}

	return dummy.Next
}

// Функция для получения k-го узла в списке
func getKthNode(head *ListNode, k int) *ListNode {
	for i := 1; i < k && head != nil; i++ {
		head = head.Next
	}
	return head
}

// Функция для переворачивания списка
func reverseList(head *ListNode, end *ListNode) {
	prev := head
	curr := head.Next
	first := curr

	for curr != end {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	head.Next = prev
	first.Next = curr
}
