package main

func middleNodeV3(head *ListNode) *ListNode {
	/*
		METHOD: Counting the Nodes
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		Суть метода:
		В этом коде мы сначала подсчитываем количество узлов в списке. Затем мы проходим по списку еще раз,
		но на этот раз мы останавливаемся на узле, который находится на позиции n/2. Это и будет наша середина.


		Time Complexity.
		Временная сложность для этого алгоритма составляет O(n), где n - количество узлов в списке.
		Это связано с тем, что мы дважды проходим по всему списку: один раз для подсчета узлов, а второй раз для поиска середины

		Space Complexity:
		Пространственная сложность для этого алгоритма составляет O(1), так как мы не используем дополнительное пространство,
		которое зависит от размера входных данных. Мы используем только некоторые переменные для подсчета узлов и для прохода по списку,
		но это не зависит от размера входных данных.
	*/
	// Инициализируем счетчик
	n, cur := 0, head

	// Пока cur не равен nil
	for cur != nil {
		// Увеличиваем счетчик
		n++
		// Переходим к следующему узлу
		cur = cur.Next
	}

	// Возвращаем серединный узел
	// Счетчик n равен количеству узлов в списке
	// Так как индексы в Go начинаются с 0, то серединный узел будет на позиции n/2
	cur = head
	for i := 0; i < n/2; i++ {
		cur = cur.Next
	}
	return cur
}
