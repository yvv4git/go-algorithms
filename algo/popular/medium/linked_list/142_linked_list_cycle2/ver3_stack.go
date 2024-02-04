package _42_linked_list_cycle2

// Функция detectCycleV3 определяет начало цикла в связном списке с помощью метода "Стек".
// Если цикла нет, то возвращается nil.
func detectCycleV3(head *ListNode) *ListNode {
	/*
		METHOD: Stack
		Time complexity: O(n^2)
		Space complexity: O(n)

		Time complexity
		Временная сложность составляет O(n^2), так как в худшем случае мы проходим по всему списку и для каждого узла мы проверяем, есть ли он в стеке.

		Space complexity
		Пространственная сложность составляет O(n), так как в худшем случае мы храним все узлы списка в стеке.
	*/

	// Используем стек для отслеживания узлов, которые мы посетили.
	stack := make([]*ListNode, 0)

	// Проходим по списку.
	for head != nil {
		// Если узел уже есть в стеке, то это начало цикла.
		for _, node := range stack {
			if node == head {
				return head
			}
		}

		// Добавляем узел в стек.
		stack = append(stack, head)

		// Переходим к следующему узлу.
		head = head.Next
	}

	// Если мы дошли до конца списка, то цикла нет.
	return nil
}
