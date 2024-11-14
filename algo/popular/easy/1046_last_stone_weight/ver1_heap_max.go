package main

import (
	"container/heap"
	"fmt"
)

// Реализуем интерфейс кучи для максимальной кучи
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // Для максимальной кучи
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	/*
		METHOD: HEAP
		Используем максимальную кучу (max-heap) для эффективного извлечения двух самых тяжелых камней.
		Куча позволяет нам извлекать максимальный элемент за O(log n) времени, что делает алгоритм эффективным.
		Max Heap (максимальная куча) — это бинарная куча, в которой каждый родительский узел имеет значение,
		которое больше или равно значениям его дочерних узлов. Это свойство называется свойством max-heap.
			9
			/ \
			7   5
			/ \ / \
			3  2 1  4

		TIME COMPLEXITY: O(n log n)
		- Построение кучи: O(n log n), так как каждый элемент добавляется в кучу за O(log n) времени.
		- Извлечение и добавление элементов: O(n log n), так как каждая операция с кучей занимает O(log n) времени.

		SPACE COMPLEXITY: O(n)
		- O(n), так как мы храним все камни в куче.
	*/
	// Инициализируем максимальную кучу
	h := &IntHeap{}
	heap.Init(h)

	// Добавляем все камни в кучу. O(n log n)
	for _, stone := range stones {
		heap.Push(h, stone) // Операция добавления элемента в кучу (heap) имеет временную сложность O(log n), а не O(n log n).
	}

	// Пока в куче больше одного камня
	for h.Len() > 1 {
		// Извлекаем два самых тяжелых камня. O(log n)
		stone1 := heap.Pop(h).(int) // Операция извлечения элемента из кучи (heap) имеет временную сложность O(log n).
		stone2 := heap.Pop(h).(int) // Операция извлечения элемента из кучи (heap) имеет временную сложность O(log n).

		// Если вес камней разный, добавляем разницу обратно в кучу. O(log n)
		if stone1 != stone2 {
			heap.Push(h, stone1-stone2) // Операция добавления элемента в кучу (heap) имеет временную сложность O(log n).
		}
	}

	// Если в куче остался один камень, возвращаем его вес
	if h.Len() == 1 {
		return (*h)[0]
	}

	// Если все камни разрушились, возвращаем 0
	return 0
}

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeight(stones)) // Вывод: 1
}
