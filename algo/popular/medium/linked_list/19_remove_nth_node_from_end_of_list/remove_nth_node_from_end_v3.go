package main

func removeNthFromEndV3(head *ListNode, n int) *ListNode {
	/*
		METHOD: Fast and slow pointers
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		В данном алгоритме два указателя fast и slow используются для создания "отступа" между ними, равного n узлам. Начальное положение обоих указателей - фиктивный узел dummy.

		1. Указатель fast перемещается на n+1 узлов вперед.
		Это делается для того, чтобы создать "отступ" между fast и slow, равный n узлам.

		2. Затем оба указателя перемещаются одновременно до тех пор, пока указатель fast не достигнет конца списка.
		Это гарантирует, что когда указатель fast достигнет конца списка, указатель slow будет на n узлах от конца списка.

		3. После того, как оба указателя достигнут нужной позиции, указатель slow используется для удаления узла,
		следующего за ним.

		Наконец, возвращается голова списка.
		Таким образом, этот алгоритм позволяет удалить n-й узел с конца списка за один проход.
	*/
	dummy := &ListNode{0, head}
	fast := dummy
	slow := dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
