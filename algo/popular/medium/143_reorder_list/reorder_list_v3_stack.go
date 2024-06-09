package main

// Функция для перестановки узлов в списке так, чтобы он стал палиндромом
func reorderListV3(head *ListNode) {
	// Если список пуст или содержит только один узел, то ничего не делаем
	if head == nil || head.Next == nil {
		return
	}

	// Инициализируем стек и добавляем в него узлы списка
	stack := []*ListNode{}
	for node := head; node != nil; node = node.Next {
		stack = append(stack, node)
	}

	// Перестраиваем узлы списка так, чтобы он стал палиндромом
	for len(stack) > 0 {
		// Извлекаем узел из стека
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Если стек пуст, то устанавливаем следующий узел в nil
		if len(stack) == 0 {
			node.Next = nil
			break
		}

		// Извлекаем следующий узел из стека
		next := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Переставляем узлы так, чтобы они образовывали палиндром
		next.Next = node.Next
		node.Next = next
	}
}
