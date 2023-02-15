package queue_on_slice

// MyQueue
/*
Push: O(1) time | O(n) space
Empty: O(1) time | O(1) space
Pop: O(1) time - amortized; 0(n) time - worst-case | O(1) space
Peek: O(1) time | O(1) space
*/
type MyQueue struct {
	S1 []int
	S2 []int
}

func New() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.S1 = append(q.S1, x)
}

func (q *MyQueue) Empty() bool {
	if len(q.S1) == 0 && len(q.S2) == 0 {
		return true
	}

	return false
}

func (q *MyQueue) Pop() int {
	var e int
	if len(q.S2) == 0 {
		for len(q.S1) != 0 {
			e = q.S1[len(q.S1)-1]
			q.S1 = q.S1[:len(q.S1)-1]
			q.S2 = append(q.S2, e)
		}
	}

	e = q.S2[len(q.S2)-1]
	q.S2 = q.S2[:len(q.S2)-1]
	return e
}

func (q *MyQueue) Peek() int {
	if len(q.S2) == 0 {
		for len(q.S1) != 0 {
			e := q.S1[len(q.S1)-1]
			q.S1 = q.S1[:len(q.S1)-1]
			q.S2 = append(q.S2, e)
		}
	}
	return q.S2[len(q.S2)-1]
}
