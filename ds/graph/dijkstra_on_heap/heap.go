/**
## ОПЕРАЦИИ
- Найти максимум или найти минимум - O(1)
- Удалить max или min элемент - O(log n)
- Увеличить или уменьшить ключ: обновить ключ - O(log n)
- Добавить новый элемент в кучу O(log n)
- Слияние двух куч в новую, которая будет содержать элементы исходных - O(n)
*/
package main

import hp "container/heap"

// Path - represents the implementation of the path.
type Path struct {
	value int
	nodes []string
}

// MinPath - used to store the minimum path.
type MinPath []Path

// Len - used for get len.
func (h MinPath) Len() int {
	return len(h)
}

// Less - used for compare values.
func (h MinPath) Less(i, j int) bool {
	return h[i].value < h[j].value
}

// Swap - used for swap values.
func (h MinPath) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push - used for push values.
func (h *MinPath) Push(x interface{}) {
	*h = append(*h, x.(Path))
}

// Pop - used for pop value.
func (h *MinPath) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Heap - graph is stored in the heap.
type Heap struct {
	values *MinPath
}

// NewHeap - used for create instance of the heap.
func NewHeap() *Heap {
	return &Heap{values: &MinPath{}}
}

// Push - used for push value.
func (h *Heap) Push(p Path) {
	hp.Push(h.values, p)
}

// Pop - used for pop value.
func (h *Heap) Pop() Path {
	i := hp.Pop(h.values)
	return i.(Path)
}
