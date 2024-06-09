package main

// CopyRandomListV2 копирует связный список с указателем на случайный узел.
func copyRandomListV2(head *Node) *Node {
	/*
		METHOD: Iterative
		Этот алгоритм основан на следующем шаблоне:
		1. Создаем копии узлов и вставляем их непосредственно за каждым узлом в исходном списке.
		2. Устанавливаем случайные указатели для копий.
		3. Разделяем исходный и скопированный списки.

		TIME COMPLEXITY: O(n), где n - количество узлов в списке, так как мы проходим по каждому узлу три раза.

		SPACE COMPLEXITY: O(1), так как мы не используем дополнительную память, кроме временной, для хранения указателей.
	*/
	if head == nil {
		return nil
	}

	// Шаг 1: Создаем копии узлов и вставляем их непосредственно за каждым узлом в исходном списке.
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = &Node{Val: cur.Val, Next: next}
		cur = next
	}

	// Шаг 2: Устанавливаем случайные указатели для копий.
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// Шаг 3: Разделяем исходный и скопированный списки.
	cur = head
	copyHead := head.Next
	for cur != nil {
		copy := cur.Next
		cur.Next = copy.Next
		if copy.Next != nil {
			copy.Next = copy.Next.Next
		}
		cur = cur.Next
	}

	return copyHead
}

//
//func main() {
//	// Создаем тестовый список.
//	node1 := &Node{Val: 1}
//	node2 := &Node{Val: 2}
//	node3 := &Node{Val: 3}
//	node1.Next = node2
//	node2.Next = node3
//	node1.Random = node3
//	node2.Random = node1
//	node3.Random = node2
//
//	// Копируем список.
//	copiedList := copyRandomList(node1)
//
//	// Выводим копию списка.
//	for copiedList != nil {
//		fmt.Printf("Val: %d, Random: %v\n", copiedList.Val, copiedList.Random)
//		copiedList = copiedList.Next
//	}
//}
