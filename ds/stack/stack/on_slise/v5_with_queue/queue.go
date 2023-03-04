package main

type Queue struct {
	items []int
}

// NewQueue - used for create instance of queue.
func NewQueue() *Queue {
	return &Queue{
		items: make([]int, 0),
	}
}

// Enqueue - used for push(add) element to queue.
// Time: O(1)
// Space: O(1)
func (q *Queue) Enqueue(val int) {
	q.items = append(q.items, val)
}

// Dequeue - used for get pop element from queue.
// Time: O(1)
// Space: O(1)
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}

	front := q.Front()
	q.items = q.items[1:]

	return front
}

// Front - used for get first element.
// Time: O(1)
// Space: O(1)
func (q *Queue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.items[0]
}

// Size - used for get size of queue.
// Time: O(1)
// Space: O(1)
func (q *Queue) Size() int {
	return len(q.items)
}

// IsEmpty - used for check queue on empty.
// Time: O(1)
// Space: O(1)
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
