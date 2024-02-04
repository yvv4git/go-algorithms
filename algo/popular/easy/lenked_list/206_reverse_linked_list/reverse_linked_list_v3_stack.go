package main

// Функция для реверса списка с использованием стека
func reverseListV3(head *ListNode) *ListNode {
	/*
		METHOD: Stack
		Time complexity: O(n)
		Space complexity: O(n)

		Time complexity.
		Временная сложность этого метода составляет O(n), где n - количество узлов в списке.
		Это связано с тем, что мы проходим по всем узлам списка два раза: один раз для добавления узлов в стек,
		а второй раз для изменения указателей.

		Space complexity.
		Пространственная сложность этого метода составляет O(n), так как мы храним все узлы списка в стеке.
	*/
	// Базовый случай: если список пуст
	if head == nil {
		return nil
	}

	// Создаем стек и добавляем все узлы в него
	stack := []*ListNode{}
	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	// Извлекаем узлы из стека и создаем новый список
	newHead := stack[len(stack)-1]
	for i := len(stack) - 1; i > 0; i-- {
		stack[i].Next = stack[i-1]
	}
	stack[0].Next = nil

	// Возвращаем новую голову списка
	return newHead
}
