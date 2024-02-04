package _2_remove_duplicates_from_sorted_list2

func deleteDuplicatesV1(head *ListNode) *ListNode {
	/*
		METHOD: Two pointers & Dummy
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность O(n), где n - количество узлов в списке. Мы проходим по списку один раз.

		Space complexity:
		Пространственная сложность O(1), так как мы используем фиктивный узел и не используем дополнительное пространство,
		зависящее от размера входных данных.

		Первый указатель, prev, будет указывать на последний узел, который не является дубликатом.
		Второй указатель, curr, будет проходить по списку.
	*/

	// Создаем фиктивный узел, который указывает на начало списка
	dummy := &ListNode{0, head}

	// Инициализируем два указателя: prev и curr
	prev := dummy
	curr := head

	// Проходим по списку
	for curr != nil {
		// Проверяем, есть ли дубликаты
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
		}

		// Если curr не сдвинулась, значит, дубликатов нет
		if prev.Next == curr {
			prev = prev.Next
		} else {
			// Если сдвинулась, значит, есть дубликаты, и мы пропускаем их
			prev.Next = curr.Next
		}

		curr = curr.Next
	}

	// Возвращаем новый список
	return dummy.Next
}
