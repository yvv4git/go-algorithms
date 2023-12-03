package main

// Функция для проверки, является ли связный список палиндромом
func isPalindromeV1(head *ListNode) bool {
	/*
		Method: Fast and slow pointers
		Time complexity: O(n)
		Space complexity: O(1)

		Time complexity для этого алгоритма составляет O(n), где n - количество узлов в связном списке.
		Это связано с тем, что мы проходим по всему списку дважды: один раз для нахождения середины списка,
		а второй раз для сравнения левой и правой половин списка.

		Space complexity для этого алгоритма составляет O(1), так как мы не используем дополнительное пространство,
		основываясь только на входных данных. Мы используем только некоторые переменные для указателей,
		но это не зависит от размера входных данных.

		Для решения этой задачи можно использовать двухпутевой подход с двумя указателями.
		Один указатель будет перемещаться с постоянной скоростью, а второй - с двумя узлами.
		Когда второй указатель достигнет конца списка, первый указатель будет на середине списка.
		Затем вы можете сравнить значения узлов с двух сторон от середины.
	*/
	// Если голова списка равна nil, то список пустой и считается палиндромом
	if head == nil {
		return true
	}

	// Инициализируем два указателя: slow и fast
	// slow и fast начинают с указателя на голову списка
	slow := head
	fast := head

	// Пока fast и fast.Next не равны nil
	for fast != nil && fast.Next != nil {
		// slow перемещается на один шаг
		slow = slow.Next
		// fast перемещается на два шага
		fast = fast.Next.Next
	}

	// Если fast не равен nil, то список имеет нечетное количество узлов
	// В этом случае, slow должен быть на середине правой половины списка
	if fast != nil {
		slow = slow.Next
	}

	// Разворачиваем правую половину списка
	slow = reverse(slow)
	fast = head

	// Проверяем, является ли список палиндромом
	for slow != nil {
		if fast.Val != slow.Val {
			return false
		}
		fast = fast.Next
		slow = slow.Next
	}

	return true
}

// Функция для разворота связного списка
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
