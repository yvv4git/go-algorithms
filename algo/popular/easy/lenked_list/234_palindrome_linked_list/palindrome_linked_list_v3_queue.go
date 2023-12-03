package main

func isPalindromeV3(head *ListNode) bool {
	/*
		Method: Queue
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	// Инициализируем очередь для хранения значений узлов
	var queue []int

	// Инициализируем два указателя: fast и slow
	// Fast перемещается на два шага за каждый шаг slow
	fast, slow := head, head

	// Пока fast и fast.Next не равны nil
	for fast != nil && fast.Next != nil {
		// Добавляем значение узла в очередь
		queue = append(queue, slow.Val)
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
		// Если значение узла не равно первому элементу очереди
		if slow.Val != queue[0] {
			// То список не является палиндромом
			return false
		}
		// Удаляем первый элемент очереди
		queue = queue[1:]
		// Переходим к следующему узлу
		slow = slow.Next
	}

	// Если мы дошли до этой точки, то список является палиндромом
	return true
}
