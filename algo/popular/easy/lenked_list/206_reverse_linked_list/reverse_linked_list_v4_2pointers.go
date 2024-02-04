package main

// Функция для реверса списка с использованием двух указателей
func reverseListV4(head *ListNode) *ListNode {
	/*
		METHOD: Two pointers
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity.
		Временная сложность этого метода составляет O(n), где n - количество узлов в списке.
		Это связано с тем, что мы проходим по всем узлам списка только один раз.

		Space complexity.
		Пространственная сложность этого метода составляет O(1), что означает постоянное пространство.
		Мы не используем дополнительное пространство, которое масштабируется с ростом входного размера.
		Мы используем только несколько указателей для отслеживания текущего узла, предыдущего узла и следующего узла.
		Поэтому пространственная сложность постоянна, независимо от размера входного значения.
	*/
	// Базовый случай: если список пуст или содержит только один элемент
	if head == nil || head.Next == nil {
		return head
	}

	// Инициализируем два указателя
	var prev *ListNode = nil
	curr := head

	// Проходим по списку и меняем указатели
	for curr != nil {
		nextTemp := curr.Next // Сохраняем следующий узел
		curr.Next = prev      // Изменяем указатель текущего узла
		prev = curr           // Перемещаем указатели на шаг вперед
		curr = nextTemp
	}

	// Возвращаем новую голову списка
	return prev
}
