package main

func partition(head *ListNode, x int) *ListNode {
	/*
		METHOD: Iterative

		TIME COMPLEXITY: O(n), где n - количество узлов в списке, потому что мы проходим по каждому узлу списка ровно один раз.

		SPACE COMPLEXITY: O(1), потому что мы используем фиксированное количество дополнительной памяти для хранения указателей before и after,
		а также фиктивных узлов beforeHead и afterHead. Независимо от размера входных данных используемая дополнительная память остается постоянной.
	*/
	// Создаем два фиктивных узла, которые будут началом двух новых списков
	beforeHead := &ListNode{0, nil}
	afterHead := &ListNode{0, nil}
	before := beforeHead
	after := afterHead

	// Проходим по исходному списку
	for head != nil {
		// Если текущий узел меньше x, добавляем его в список before
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			// Иначе добавляем его в список after
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}

	// Заканчиваем список after
	after.Next = nil

	// Объединяем два списка
	before.Next = afterHead.Next

	// Возвращаем голову объединенного списка
	return beforeHead.Next
}
