package queue_using_stack

type Node struct {
	Val  int
	Next *Node
}

type MyQueue struct {
	Head *Node
	Tail *Node
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	newNode := &Node{Val: x}
	if q.Head == nil {
		q.Head = newNode
		q.Tail = newNode
		return
	}
	q.Tail.Next = newNode
	q.Tail = newNode
}

func (q *MyQueue) Pop() int {
	temp := q.Head
	q.Head = q.Head.Next
	return temp.Val
}

func (q *MyQueue) Peek() int {
	return q.Head.Val
}

func (q *MyQueue) Empty() bool {
	return q.Head == nil
}
