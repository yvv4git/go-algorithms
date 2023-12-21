package _32_implement_queue_using_stacks

/*
Time complexity
Временная сложность всех методов в данном коде - O(1), так как все операции выполняются за постоянное время, независимо от размера очереди.

Space complexity
Пространственная сложность - O(n), где n - количество узлов в очереди. Это связано с тем, что для каждого узла в очереди выделяется отдельная память.
*/

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

// Push - time complexity O(1)
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

// Pop - time complexity O(1)
func (q *MyQueue) Pop() int {
	temp := q.Head
	q.Head = q.Head.Next
	return temp.Val
}

// Peek - time complexity O(1)
func (q *MyQueue) Peek() int {
	return q.Head.Val
}

// Empty - time complexity O(1)
func (q *MyQueue) Empty() bool {
	return q.Head == nil
}
