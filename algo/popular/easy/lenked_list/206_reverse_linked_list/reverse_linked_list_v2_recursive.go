package main

// Функция для рекурсивного реверса списка
func reverseListV2(head *ListNode) *ListNode {
	/*
		Method: Recursive
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity.
		Временная сложность этого метода составляет O(n), где n - количество узлов в списке.
		Это связано с тем, что мы проходим по всем узлам списка только один раз.

		Space complexity:
		Пространственная сложность этого метода составляет O(n),
		так как в худшем случае глубина рекурсии может достигать n (когда список является связным списком с n узлами).
	*/
	// Базовый случай: если список пуст или содержит только один элемент
	if head == nil || head.Next == nil {
		return head
	}

	// Рекурсивный случай: рекурсивно реверсируем оставшуюся часть списка
	newHead := reverseListV2(head.Next)

	// Изменяем указатель текущего узла
	head.Next.Next = head
	head.Next = nil

	// Возвращаем новую голову списка
	return newHead
}
