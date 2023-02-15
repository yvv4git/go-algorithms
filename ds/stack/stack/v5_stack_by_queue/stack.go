package main

type MyStack struct {
	q1, q2 *Queue
}

// NewStack - used for create stack instance.
func NewStack() MyStack {
	return MyStack{
		q1: NewQueue(),
		q2: NewQueue(),
	}
}

// Push - used for push element to stack.
// Time: O(n + 1 + n) = O(2n+1) = O(n)
// Space: O(1)
func (s *MyStack) Push(val int) {
	// Time: O(n)
	swap(s.q1, s.q2)
	// Time: O(1)
	s.q1.Enqueue(val)
	// Time: O(n)
	swap(s.q2, s.q1)
}

// Time: O(n)
// Space: O(1)
func swap(q1, q2 *Queue) {
	for !q1.IsEmpty() {
		q2.Enqueue(q1.Dequeue())
	}
}

// Pop - used for pop element from stack.
// Time: O(1)
// Space: O(1)
func (s *MyStack) Pop() int {
	return s.q1.Dequeue()
}

// Top - used for get top element.
// Time: O(1)
// Space: O(1)
func (s *MyStack) Top() int {
	return s.q1.Front()
}

// IsEmpty - used for check is emptyStack.
// I would have named this IsEmpty()...
// Time: O(1)
// Space: O(1)
func (s *MyStack) IsEmpty() bool {
	return s.q1.IsEmpty()
}
