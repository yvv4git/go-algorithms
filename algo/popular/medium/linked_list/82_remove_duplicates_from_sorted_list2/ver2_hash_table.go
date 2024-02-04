package _2_remove_duplicates_from_sorted_list2

func deleteDuplicatesV2(head *ListNode) *ListNode {
	/*
		METHOD: Hash table, Dictionary
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	// Создаем фиктивный узел, который указывает на начало списка
	dummy := &ListNode{0, head}

	// Инициализируем два указателя: prev и curr
	prev := dummy
	curr := head

	// Создаем хэш-таблицу для отслеживания уже встречавшихся элементов
	seen := make(map[int]bool)

	// Проходим по списку
	for curr != nil {
		// Проверяем, был ли элемент уже встречавшимся
		if _, ok := seen[curr.Val]; ok {
			// Если был, то пропускаем его
			prev.Next = curr.Next
		} else {
			// Если не был, то добавляем его в хэш-таблицу и переходим к следующему узлу
			seen[curr.Val] = true
			prev = curr
		}

		curr = curr.Next
	}

	// Возвращаем новый список
	return dummy.Next
}
