package main

func deleteNode(node *ListNode) {
	/*
		METHOD: Iterative

		TIME COMPLEXITY:  O(1), так как мы выполняем фиксированное количество операций независимо от размера входных данных.
		Это связано с тем, что мы всего лишь копируем значение и перенаправляем указатель, не проходя по всему списку.

		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество дополнительной памяти для выполнения операций,
		не зависящих от размера входных данных.
	*/
	// Проверяем, что узел не является последним в списке
	if node.Next != nil {
		// Копируем значение следующего узла в текущий узел
		node.Val = node.Next.Val
		// Перенаправляем указатель текущего узла на следующий за следующим узел
		node.Next = node.Next.Next
	}
}
