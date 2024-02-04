package _42_linked_list_cycle2

// Функция detectCycleV2 определяет начало цикла в связном списке с помощью метода Hash Set.
// Если цикла нет, то возвращается nil.
func detectCycleV2(head *ListNode) *ListNode {
	/*
		METHOD: Hash set
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)

		Time complexity
		Временная сложность составляет O(n), так как в худшем случае мы проходим по всему списку.

		Space complexity
		Пространственная сложность составляет O(n), так как в худшем случае мы храним все узлы списка в Hash Set.
	*/

	// Используем Hash Set для отслеживания уже посещенных узлов.
	visited := make(map[*ListNode]bool)

	// Проходим по списку.
	for head != nil {
		// Если узел уже был посещен, то это начало цикла.
		if visited[head] {
			return head
		}

		// Помечаем узел как посещенный.
		visited[head] = true

		// Переходим к следующему узлу.
		head = head.Next
	}

	// Если мы дошли до конца списка, то цикла нет.
	return nil
}
