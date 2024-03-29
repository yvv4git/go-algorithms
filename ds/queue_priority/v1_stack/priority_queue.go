package v1_stack

import (
	"container/heap"
)

// Prque - priority queue data structure.
type Prque struct {
	cont *sstack
}

// New - creates a new priority queue.
func New(setIndex setIndexCallback) *Prque {
	return &Prque{newSstack(setIndex)}
}

// Push - pushes a value with a given priority into the queue, expanding if necessary.
func (p *Prque) Push(data interface{}, priority int64) {
	heap.Push(p.cont, &item{data, priority})
}

// Pop - pops the value with the greates priority off the stack and returns it.
func (p *Prque) Pop() (interface{}, int64) {
	item := heap.Pop(p.cont).(*item)
	return item.value, item.priority
}

// PopItem - pops only the item from the queue, dropping the associated priority value.
func (p *Prque) PopItem() interface{} {
	return heap.Pop(p.cont).(*item).value
}

// Remove removes the element with the given index.
func (p *Prque) Remove(i int) interface{} {
	if i < 0 {
		return nil
	}
	return heap.Remove(p.cont, i)
}

// Empty - checks whether the priority queue is empty.
func (p *Prque) Empty() bool {
	return p.cont.Len() == 0
}

// Size - returns the number of element in the priority queue.
func (p *Prque) Size() int {
	return p.cont.Len()
}

// Reset - clears the contents of the priority queue.
func (p *Prque) Reset() {
	*p = *New(p.cont.setIndex)
}
