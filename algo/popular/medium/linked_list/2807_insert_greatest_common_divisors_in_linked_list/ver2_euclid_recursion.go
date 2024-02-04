package _807_insert_greatest_common_divisors_in_linked_list

func insertGreatestCommonDivisorsV2(head *ListNode) *ListNode {
	/*
		METHOD: Euclid + recursion
		Time complexity: O(n)
		Space complexity:  O(n)

		Time complexity:
		Временная сложность этого алгоритма - O(n), где n - количество узлов в списке.
		Это связано с тем, что мы проходим по каждому узлу списка ровно один раз.

		Space complexity:
		Пространственная сложность - O(n), так как в худшем случае (когда мы добавляем новый узел между каждыми двумя узлами) мы используем O(n) дополнительной памяти.
	*/
	// Базовый случай: если список содержит менее двух узлов, то ничего не делаем
	if head == nil || head.Next == nil {
		return head
	}

	// Создаем новый узел с НОД двух первых узлов списка
	newNode := &ListNode{Val: gcd(head.Val, head.Next.Val)}
	newNode.Next = head.Next
	head.Next = newNode

	// Рекурсивный вызов функции для следующего узла
	insertGreatestCommonDivisorsV2(newNode.Next.Next)

	return head
}
