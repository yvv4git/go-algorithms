package main

func reverseListV1(head *ListNode) *ListNode {
	/*
		METHOD: Iterative Reversal
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		TIME COMPLEXITY:
		Временная сложность этого метода составляет O(n), где n - количество узлов в списке.
		Это связано с тем, что мы проходим по всем узлам списка только один раз.

		SPACE COMPLEXITY:
		Пространственная сложность этого метода составляет O(1), что означает постоянное пространство.
		Мы не используем дополнительное пространство, которое масштабируется с ростом входного размера.
		Мы используем только несколько указателей для отслеживания текущего узла, предыдущего узла и следующего узла.
		Поэтому пространственная сложность постоянна, независимо от размера входного значения.
	*/
	var prev *ListNode
	for head != nil {
		nextNode := head.Next
		head.Next = prev
		prev = head
		head = nextNode
	}
	return prev
}
