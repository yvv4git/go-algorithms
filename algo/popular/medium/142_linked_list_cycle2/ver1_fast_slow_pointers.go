package _42_linked_list_cycle2

// Функция detectCycleV1 определяет начало цикла в связном списке.
// Если цикла нет, то возвращается nil.
func detectCycleV1(head *ListNode) *ListNode {
	/*
		METHOD: Floyd's Tortoise and Hare / Fast and slow pointers
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)

		Time complexity
		Временная сложность составляет O(n), так как в худшем случае мы проходим по всему списку два раза:
		один раз для определения наличия цикла, а второй раз для определения начала цикла.

		Space complexity
		Пространственная сложность составляет O(1), так как алгоритм использует постоянное количество памяти,
		независимо от размера входных данных.

		Метод, используемый в этом коде, называется "Floyd's Tortoise and Hare".
		Это алгоритм, который используется для поиска циклов в связных списках.
	*/

	// Если список пуст или содержит только один элемент, то в нем не может быть цикла.
	if head == nil || head.Next == nil {
		return nil
	}

	// Используем два указателя - быстрый и медленный.
	slow, fast := head, head

	// Ищем цикл.
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		// Если быстрый и медленный указатели встретились, то цикл найден.
		if slow == fast {
			break
		}
	}

	// Если быстрый указатель дошел до конца списка, то цикла нет.
	if fast == nil || fast.Next == nil {
		return nil
	}

	// Если цикл найден, то определяем его начало.
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	// Возвращаем начало цикла.
	return slow
}
