package main

import (
	"container/heap"
)

// Heap представляет собой максимальную кучу (max-heap)
type Heap []int

// Len возвращает длину кучи
func (h Heap) Len() int { return len(h) }

// Less определяет порядок элементов в куче (для max-heap возвращает true, если i-й элемент больше j-го)
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }

// Swap меняет местами i-й и j-й элементы в куче
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// Push добавляет элемент в кучу
func (h *Heap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

// Pop удаляет и возвращает максимальный элемент из кучи
func (h *Heap) Pop() interface{} {
	cur := *h
	e := cur[len(cur)-1]
	*h = cur[:len(cur)-1]
	return e
}

func kthSmallestV4(matrix [][]int, k int) int {
	/*
		METHOD: Max-Heap (Максимальная куча)
		DESCRIPTION: Используется максимальная куча для нахождения k-го наименьшего элемента в матрице.
		Куча поддерживает не более k элементов, что позволяет эффективно находить k-й наименьший элемент.

		TIME COMPLEXITY: O(n * m * log(k))
		DESCRIPTION: Временная сложность определяется проходом по всем элементам матрицы (n * m) и
		операциями вставки и удаления в куче размера k (log(k)).

		SPACE COMPLEXITY: O(k)
		DESCRIPTION: Пространственная сложность определяется размером кучи, которая хранит не более k элементов.
	*/
	// Инициализируем максимальную кучу
	maxHeap := &Heap{}
	heap.Init(maxHeap)

	// Проходим по всем элементам матрицы
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// Добавляем элемент в кучу
			heap.Push(maxHeap, matrix[i][j])
			// Если размер кучи превышает k, удаляем максимальный элемент
			if len(*maxHeap) > k {
				heap.Pop(maxHeap)
			}
		}
	}

	// Если размер кучи меньше k, возвращаем -1 (ошибка)
	if len(*maxHeap) < k {
		return -1
	}

	// Возвращаем k-й наименьший элемент (максимальный элемент в куче)
	return heap.Pop(maxHeap).(int)
}
