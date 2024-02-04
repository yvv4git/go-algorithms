package _3_merge_k_sorted_lists

import "container/heap"

func mergeKListsV1(lists []*ListNode) *ListNode {
	/*
		METHOD: Priority Queue
		TIME COMPLEXITY: O(N log k), N - общее количество узлов во всех списках, k - количество списков
		Space complexity: O(k), k - количество списков

		TIME COMPLEXITY:
		Временная сложность этого алгоритма составляет O(N log k), где N - общее количество узлов во всех списках, k - количество списков.
		Это связано с тем, что мы вставляем и извлекаем каждый узел из очереди ровно один раз.

		Space complexity:
		Пространственная сложность составляет O(k), так как в любой момент времени в очереди будет храниться не более k узлов.


		Под капотом используется Priority Queue, которая позволяет извлекать минимальный элемент за O(1) времени.
		В этом случае, минимальным элементом является узел с минимальным значением в списке.

		В начале алгоритма все первые элементы из всех списков помещаются в очередь.
		Затем из очереди извлекаются элементы, которые меньше всех остальных, и добавляются в результирующий список.
		Если у извлеченного элемента есть следующий элемент в списке, он также помещается в очередь.

		Этот процесс повторяется до тех пор, пока очередь не станет пустой.
	*/
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Добавляем первый элемент из каждого списка в очередь
	for _, list := range lists {
		if list != nil {
			heap.Push(pq, list)
		}
	}

	// Создаем фиктивный узел, который будет началом результирующего списка
	dummy := &ListNode{}
	curr := dummy

	// Извлекаем элементы из очереди и добавляем их в результирующий список
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*ListNode)
		curr.Next = node
		curr = curr.Next

		if node.Next != nil {
			heap.Push(pq, node.Next)
		}
	}

	return dummy.Next
}

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// PriorityQueue - Решение будет использовать приоритетную очередь (heap)
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Здесь мы используем оператор "<" для создания "min heap"
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return x
}
