package main

// Функция mergeTwoListsV1 объединяет два отсортированных списка в один отсортированный список.
// Входные данные: l1 и l2 - указатели на начало списков.
// Выходные данные: указатель на начало объединенного списка.
func mergeTwoListsV1(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		METHOD: Dummy node
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		Временная сложность: O(n), где n - количество узлов в обоих списках.
		Это происходит, потому что мы проходимся по каждому узлу в обоих списках по одному разу.

		Пространственная сложность: O(1), поскольку мы используем фиктивный узел dummy,
		который не учитывается в пространственной сложности.
	*/
	// Создаем фиктивный узел, который будет началом результирующего списка.
	dummy := ListNode{}
	dummy.Next = nil
	// Указатель tail используется для отслеживания последнего узла в результирующем списке.
	tail := &dummy

	// Проходимся по обоим спискам, сравнивая элементы и добавляя меньший из них в результирующий список.
	for l1 != nil || l2 != nil {
		if l1 == nil {
			// Если первый список закончился, добавляем оставшиеся элементы из второго списка.
			(*tail).Next = l2
			break
		} else if l2 == nil {
			// Если второй список закончился, добавляем оставшиеся элементы из первого списка.
			(*tail).Next = l1
			break
		} else {
			// Сравниваем элементы из обоих списков и добавляем меньший в результирующий список.
			if l1.Val <= l2.Val {
				(*tail).Next = l1
				l1 = (*l1).Next
			} else {
				(*tail).Next = l2
				l2 = (*l2).Next
			}
			// Перемещаем указатель tail на следующий узел.
			tail = tail.Next
		}
	}

	// Возвращаем указатель на начало объединенного списка.
	return dummy.Next
}
