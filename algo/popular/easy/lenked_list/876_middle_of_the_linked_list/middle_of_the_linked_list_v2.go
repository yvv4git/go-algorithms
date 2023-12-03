package main

// Функция для нахождения середины списка
func middleNodeV2(head *ListNode) *ListNode {
	// Инициализируем стек
	var stack []*ListNode

	// Пока head не равен nil
	for head != nil {
		// Добавляем в стек текущий узел
		stack = append(stack, head)
		// Переходим к следующему узлу
		head = head.Next
	}

	// Возвращаем серединный узел
	// Длина стека stack равна количеству узлов в списке
	// Так как индексы в Go начинаются с 0, то серединный узел будет на позиции len(stack)/2
	return stack[len(stack)/2]
}
