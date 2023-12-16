package _5_reverse_nodes_in_k_group

// Функция reverseKGroup разворачивает каждую группу k узлов в связном списке.
// Она возвращает новый начальный узел после разворота.
// Временная сложность: O(n), где n - количество узлов в списке.
// Пространственная сложность: O(1).
func reverseKGroupV3(head *ListNode, k int) *ListNode {
	/*
		Method: Two pointer
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity
		Временная сложность каждой функции в худшем случае составляет O(n), где n - количество узлов в списке.
		Это происходит потому, что функции проходят по всем узлам в списке только один раз.

		Space complexity
		Пространственная сложность каждой функции также составляет O(1),
		так как функции используют фиксированное количество переменных, не зависящих от размера входного списка.

		Этот метод использует два указателя, обычно называемые first и last, для обхода списка.
		Один указатель начинает с начала списка, а второй указатель начинает с k-го узла.
		Затем меняются местами узлы между first и last, пока first не достигнет last.
	*/
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head // Недостаточно узлов для разворота группы
		}
		node = node.Next
	}

	newHead := reverseV2(head, node)
	head.Next = reverseKGroupV3(node, k)
	return newHead
}

// Функция reverse меняет местами узлы в связном списке от first до last.
// Она возвращает новый начальный узел после разворота.
// Временная сложность: O(n), где n - количество узлов в диапазоне от first до last.
// Пространственная сложность: O(1).
func reverseV2(first *ListNode, last *ListNode) *ListNode {
	prev := last
	for first != last {
		next := first.Next
		first.Next = prev
		prev = first
		first = next
	}
	return prev
}
