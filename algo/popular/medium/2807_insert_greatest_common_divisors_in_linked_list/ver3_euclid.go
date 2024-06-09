package _807_insert_greatest_common_divisors_in_linked_list

func insertGreatestCommonDivisorsV3(head *ListNode) *ListNode {
	/*
		METHOD:
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY:	O(1)

		Данный метод insertGreatestCommonDivisors использует итеративный подход для добавления новых узлов в связный список.
		Каждый новый узел содержит наибольший общий делитель (НОД) двух соседних узлов.
	*/
	res, prev, head := head, head, head.Next
	for head != nil { // Цикл, который продолжается до тех пор, пока head не станет равным nil, что означает, что мы достигли конца списка
		prev.Next = &ListNode{Val: gcd(prev.Val, head.Val), Next: head} // Создание нового узла, который содержит НОД двух соседних узлов. Этот новый узел становится следующим узлом для prev
		prev = head
		head = head.Next
	}

	// Возвращение указателя на первый узел списка
	return res
}
