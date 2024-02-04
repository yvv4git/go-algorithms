package _3_merge_k_sorted_lists

// MergeKListsV2 - функция, которая объединяет k отсортированных связанных списков в один отсортированный список.
// Она использует итеративный подход.
// Временная сложность: O(N*M), где N - общее количество узлов, M - количество списков.
// Пространственная сложность: O(1), так как дополнительной памяти не используется.
func mergeKListsV2(lists []*ListNode) *ListNode {
	/*
		METHOD: Iterative
		TIME COMPLEXITY: O(N*M), где N - общее количество узлов, M - количество списков
		SPACE COMPLEXITY: O(1), так как дополнительной памяти не используется
	*/
	// Если список пуст, возвращаем nil
	if len(lists) == 0 {
		return nil
	}

	// Устанавливаем голову объединенного списка в первый список
	var head *ListNode = lists[0]

	// Если есть только один список, возвращаем его
	if len(lists) == 1 {
		return head
	}

	// Итерируемся по оставшимся спискам и объединяем их с головой списка
	for _, list := range lists[1:] {
		head = mergeTwoLists(head, list)
	}

	// Возвращаем голову объединенного списка
	return head
}

func mergeTwoLists(list1, list2 *ListNode) (head *ListNode) {
	list1, list2, head = removeSmallerNode(list1, list2)

	current := head
	for list1 != nil && list2 != nil {
		var next *ListNode
		list1, list2, next = removeSmallerNode(list1, list2)
		current.Next, current = next, next
	}
	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return head
}

func removeSmallerNode(list1, list2 *ListNode) (l1, l2, node *ListNode) {
	if list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			node, list1 = list1, list1.Next
		} else {
			node, list2 = list2, list2.Next
		}
	} else if list1 != nil /* && list2 == nil */ {
		node, list1 = list1, list1.Next
	} else if list2 != nil /* && list1 == nil */ {
		node, list2 = list2, list2.Next
	}
	return list1, list2, node
}
