package _807_insert_greatest_common_divisors_in_linked_list

// Функция для вставки НОД между узлами
func insertGreatestCommonDivisorsV1(head *ListNode) *ListNode {
	/*
		METHOD:
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)

		TIME COMPLEXITY:
		Временная сложность этого алгоритма составляет O(n), где n - количество узлов в связанном списке.
		Это связано с тем, что мы проходим по каждому узлу списка ровно один раз.

		Space complexity:
		Пространственная сложность также составляет O(n), так как в худшем случае мы можем хранить в памяти все узлы списка.
	*/
	curr := head
	for curr != nil && curr.Next != nil {
		gcdVal := gcd(curr.Val, curr.Next.Val) // Вычисляем НОД двух значений текущего и следующего узлов
		newNode := &ListNode{Val: gcdVal}      // Создаем новый узел с вычисленным НОД
		newNode.Next = curr.Next               // Вставляем новый узел между текущим и следующим узлами
		curr.Next = newNode

		// Переходим к следующей паре узлов
		curr = curr.Next.Next
	}
	return head
}
