package main

// Функция для проверки, является ли список палиндромом
func isPalindromeV2(head *ListNode) bool {
	/*
		METHOD: Stack
		Time Complexity: O(n)
		Space Complexity: O(n)

		Time Complexity:
		Временная сложность для этого алгоритма составляет O(n), где n - количество узлов в списке.
		Это связано с тем, что мы проходим по всему списку дважды: один раз для добавления значений в стек,
		а второй раз для сравнения значений.

		Space Complexity:
		Пространственная сложность для этого алгоритма составляет O(n), так как мы создаем новый стек для хранения значений узлов.
		Размер стека зависит от количества узлов в списке.
	*/
	// Инициализируем стек для хранения значений узлов
	var stack []int

	// Инициализируем два указателя: fast и slow
	// Fast перемещается на два шага за каждый шаг slow
	fast, slow := head, head

	// Пока fast и fast.Next не равны nil
	for fast != nil && fast.Next != nil {
		// Добавляем значение узла в стек
		stack = append(stack, slow.Val)
		// Переходим к следующему узлу
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Если длина списка - нечетное число, то пропускаем серединный элемент
	if fast != nil {
		slow = slow.Next
	}

	// Пока slow не равен nil
	for slow != nil {
		// Если значение узла не равно последнему элементу стека
		if slow.Val != stack[len(stack)-1] {
			// То список не является палиндромом
			return false
		}
		// Удаляем последний элемент стека
		stack = stack[:len(stack)-1]
		// Переходим к следующему узлу
		slow = slow.Next
	}

	// Если мы дошли до этой точки, то список является палиндромом
	return true
}
