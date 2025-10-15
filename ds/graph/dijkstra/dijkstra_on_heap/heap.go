package dijkstra

import "container/heap"

// Heap - graph is stored in the heap.
type Heap struct {
	MinPath *MinPath
}

// NewHeap - used for create instance of the heap.
func NewHeap() *Heap {
	return &Heap{
		MinPath: &MinPath{},
	}
}

// Push - used for push weight.
func (h *Heap) Push(p Path) {
	heap.Push(h.MinPath, p)
}

// Pop - used for pop weight.
func (h *Heap) Pop() Path {
	i := heap.Pop(h.MinPath)
	return i.(Path)
}
