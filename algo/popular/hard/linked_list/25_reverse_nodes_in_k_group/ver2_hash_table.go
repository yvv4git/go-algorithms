package _5_reverse_nodes_in_k_group

func reverseKGroupV2(head *ListNode, k int) *ListNode {
	/*
		METHOD: Hash table, Dictionary
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность функции reverse в худшем случае составляет O(n), где n - количество узлов в диапазоне от first до last.
		Это происходит потому, что функция проходит по всем узлам в диапазоне от first до last только один раз.

		Space complexity:
		Пространственная сложность функции reverse также составляет O(1),
		так как функция использует фиксированное количество переменных, не зависящих от размера входного списка.
	*/
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head // если количество узлов меньше k, то мы просто возвращаем исходный список
		}
		node = node.Next
	}
	newHead := reverse(head, node)       // меняем местами узлы в группе
	head.Next = reverseKGroupV2(node, k) // рекурсивно меняем местами узлы в следующих группах
	return newHead
}

// Функция reverse меняет местами узлы в связном списке от first до last.
// Она возвращает новый начальный узел после разворота.
func reverse(first *ListNode, last *ListNode) *ListNode {
	// prev - это указатель на предыдущий узел. Изначально он указывает на последний узел.
	prev := last

	// Цикл продолжается, пока first не станет равным last.
	// Это означает, что мы пройдем по всем узлам в диапазоне от first до last.
	for first != last {
		// next - это указатель на следующий узел после first.
		next := first.Next

		// Меняем местами указатели следующего узла для first и prev.
		// Это означает, что первый узел будет указывать на предыдущий узел, а не на следующий.
		first.Next = prev

		// Перемещаем указатели prev и first на один узел вперед.
		prev = first
		first = next
	}

	// Возвращаем новый начальный узел после разворота.
	return prev
}
