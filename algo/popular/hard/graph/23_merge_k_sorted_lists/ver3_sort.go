package _3_merge_k_sorted_lists

import "sort"

// Время: O(n + nlogn + n) = O(nlogn + 2n) = O(nlogn + n) = O(nlogn)
// Пространственная сложность: O(n)
func mergeKListsV3(lists []*ListNode) *ListNode {
	/*
		METHOD: Sort
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: O(n)

		Что тут происходит:
		1. Деление.
		В методе mergeKLists список списков lists делится на отдельные списки. Каждый список преобразуется в отдельный список узлов с помощью функции convertToNodes.

		2. Власть.
		Каждый отдельный список узлов сортируется с помощью функции sortNodes.

		3. Объединение.
		После сортировки, все отсортированные списки объединяются в один отсортированный список узлов с помощью функции convertToList
	*/
	// Время: O(n)
	// Пространственная сложность: O(n)
	nodes := convertToNodes(lists)
	// Время: O(nlogn)
	sortNodes(nodes)
	// Время: O(n)
	return convertToList(nodes)
}

// merge преобразует срез связанных списков в один связанный список и возвращает его голову.
// Время: O(n)
// Пространственная сложность: O(n) (Возвращаемые данные)
// Где `n` - количество узлов.
func convertToList(nodes []*ListNode) (head *ListNode) {
	if len(nodes) == 0 {
		return nil
	}

	head = nodes[0]
	// Время: O(1)
	if current := head; current != nil && len(nodes) > 1 {
		// Время: O(n)
		for _, node := range nodes[1:] {
			current.Next, current = node, node
		}
	}
	return head
}

// sort сортирует срез связанных списков.
// Все узлы среза должны не быть частью связанного списка (они все указывают на null).
// Время: O(nlogn)
// Пространственная сложность: O(1)
// Где `n` - количество узлов.
func sortNodes(nodes []*ListNode) {
	// Время: O(nlogn)
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})
}

// merge преобразует все связанные списки среза
// в срез связанных списков узлов
// и объединяет их все.
// Время: O(n)
// Пространственная сложность: O(n) (Возвращаемые данные)
// Где
// - `n`: общее количество узлов n=x*k
// - `x` - длина списков
// - `k` - среднее количество узлов.
func convertToNodes(lists []*ListNode) (nodes []*ListNode) {
	// Время: O(x)
	for _, list := range lists {
		// Время: O(k) - среднее количество узлов
		nodes = append(nodes, convertNode(list)...)
	}
	return nodes
}

// convert преобразует связанный список
// в срез связанных списков узлов
// (все они указывают на null).
// Время: O(k)
// Пространственная сложность: O(n) (Возвращаемые данные)
// Где `k` - количество узлов.
func convertNode(node *ListNode) (nodes []*ListNode) {
	// Время: O(k)
	for node != nil {
		next := node.Next
		node.Next = nil
		nodes = append(nodes, node)
		node = next
	}
	return nodes
}
