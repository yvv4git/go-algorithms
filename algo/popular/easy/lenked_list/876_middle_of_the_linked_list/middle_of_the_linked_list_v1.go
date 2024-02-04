package main

// Функция для нахождения середины связного списка
func middleNodeV1(head *ListNode) *ListNode {
	/*
		METHOD: Fast and slow pointers
		Time complexity: O(n)
		Space complexity: O(1)

		Особенность метода Fast & slow pointers!
		Когда fast указатель достигает конца списка, это означает, что список имеет четное количество узлов.
		В этом случае, когда fast достигает конца списка, slow будет указывать на второй узел середины списка.
		Если список имеет нечетное количество узлов, то fast достигнет конца, когда fast.Next будет nil.
		В этом случае, slow будет указывать на середину списка.
	*/

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

	// Когда fast достигнет конца списка, slow будет на середине списка
	return slow
}
