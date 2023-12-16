package _2_remove_duplicates_from_sorted_list2

func deleteDuplicatesV4(head *ListNode) *ListNode {
	/*
		Method: Stack
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность O(n), где n - количество узлов в списке. Мы проходим по списку один раз.

		Space complexity
		Пространственная сложность: O(n), так как в стеке мы храним все уникальные значения узлов.
		В худшем случае, когда все значения уникальны, мы будем хранить n элементов.
	*/
	// Создаем фиктивный узел, который указывает на начало списка
	dummy := &ListNode{0, head}

	// Инициализируем два указателя: prev и curr
	prev := dummy
	curr := head

	// Создаем стек для отслеживания уже встречавшихся элементов
	stack := []int{}

	// Проходим по списку
	for curr != nil {
		// Проверяем, был ли элемент уже встречавшимся
		if len(stack) > 0 && stack[len(stack)-1] == curr.Val {
			// Если был, то пропускаем его
			for curr != nil && stack[len(stack)-1] == curr.Val {
				curr = curr.Next
			}
			// Удаляем последний элемент из стека
			stack = stack[:len(stack)-1]
			// Устанавливаем следующий узел для предыдущего узла
			if len(stack) > 0 {
				prev.Next = curr
			} else {
				prev.Next = nil
			}
		} else {
			// Если не был, то добавляем его в стек и переходим к следующему узлу
			stack = append(stack, curr.Val)
			prev = curr
			curr = curr.Next
		}
	}

	// Возвращаем новый список
	return dummy.Next
}
