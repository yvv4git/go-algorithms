package _4_swap_nodes_in_pairs

// Функция для перестановки узлов в парах
func swapPairsV1(head *ListNode) *ListNode {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/

	// Создаем фиктивный узел
	dummy := &ListNode{0, head}

	// Инициализируем указатели prev и curr
	prev, curr := dummy, head

	// Проходим по списку
	for curr != nil && curr.Next != nil {
		// Сохраняем указатели на следующие узлы
		nextPair := curr.Next.Next
		nextNode := curr.Next

		// Меняем указатели на следующие узлы
		nextNode.Next = curr
		curr.Next = nextPair
		prev.Next = nextNode

		// Перемещаем указатели
		prev = curr
		curr = nextPair
	}

	// Возвращаем новую голову списка
	return dummy.Next
}
