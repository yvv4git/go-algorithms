package _2_remove_duplicates_from_sorted_list2

func deleteDuplicatesV3(head *ListNode) *ListNode {
	/*
		METHOD: Recursion
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность O(n), где n - количество узлов в списке. Мы проходим по списку один раз.

		Space complexity
		Пространственная сложность O(n), так как в худшем случае, когда все значения уникальны,
		мы можем иметь n вызовов рекурсии на стеке.
	*/
	// Базовый случай: если список пуст или содержит только один элемент
	if head == nil || head.Next == nil {
		return head
	}

	// Рекурсивный случай: если следующий узел имеет то же значение, что и текущий
	if head.Next.Val == head.Val {
		// Пропускаем все следующие узлы с таким же значением
		for head.Next != nil && head.Next.Val == head.Val {
			head = head.Next
		}
		// Пропускаем текущий узел и все следующие узлы с таким же значением
		return deleteDuplicatesV3(head.Next)
	} else {
		// Если следующий узел имеет другое значение, то проверяем следующий узел
		head.Next = deleteDuplicatesV3(head.Next)
	}

	return head
}
