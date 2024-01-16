package main

import "container/heap"

// KthLargest - структура для хранения k наибольших элементов
type KthLargest struct {
	k    int
	heap *IntHeap
}

// Constructor - конструктор для KthLargest
func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return KthLargest{k: k, heap: h}
}

// Add - добавляет новый элемент в кучу
func (this *KthLargest) Add(val int) int {
	/*
		Time complexity: O(log k), так как вставка элемента в кучу занимает O(log k) времени.
		Space complexity: O(k), так как в худшем случае мы храним k элементов в куче.

		Метод Add добавляет новый элемент в кучу, а затем, если размер кучи превышает k, удаляет наименьший элемент из кучи.
		Время выполнения этих операций зависит от размера кучи, поэтому временная сложность составляет O(log k).
		Пространственная сложность составляет O(k), так как в худшем случае мы храним k элементов в куче.
	*/
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}

	return (*this.heap)[0]
}

// IntHeap - куча для целых чисел
type IntHeap []int

// Len возвращает количество элементов в куче.
// Временная сложность: O(1)
func (h IntHeap) Len() int { return len(h) }

// Less сравнивает два элемента в куче.
// Временная сложность: O(1)
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

// Swap меняет местами два элемента в куче.
// Временная сложность: O(1)
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push добавляет новый элемент в кучу.
// Временная сложность: O(1)
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop удаляет наименьший элемент из кучи.
// Временная сложность: O(1)
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
