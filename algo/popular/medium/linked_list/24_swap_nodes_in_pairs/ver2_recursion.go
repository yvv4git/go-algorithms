package _4_swap_nodes_in_pairs

// Функция для перестановки узлов в парах
func swapPairsV2(head *ListNode) *ListNode {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)

		TIME COMPLEXITY:
		Временная сложность этого алгоритма составляет O(n), где n - количество узлов в списке, потому что мы проходим по списку один раз.

		Space complexity:
		Пространственная сложность составляет O(n), поскольку мы используем рекурсию, которая может добавить дополнительные уровни рекурсии в стеке вызовов.
	*/

	// Базовый случай: если список пуст или содержит только один узел
	if head == nil || head.Next == nil {
		return head
	}

	// Получаем следующий узел
	next := head.Next

	// Получаем узел после следующего узла
	nextNext := next.Next

	// Меняем указатели
	next.Next = head
	head.Next = swapPairsV2(nextNext)

	// Возвращаем новый головной узел
	return next
}
