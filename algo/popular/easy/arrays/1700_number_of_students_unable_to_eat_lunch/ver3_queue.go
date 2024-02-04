package _700_number_of_students_unable_to_eat_lunch

type IntNode struct {
	prev *IntNode
	val  int
	next *IntNode
}

type IntQueue struct {
	head   *IntNode
	tail   *IntNode
	Length int
}

// NewIntQueue создает новую очередь
func NewIntQueue() *IntQueue {
	q := IntQueue{}

	q.head = &IntNode{val: 0}
	q.tail = &IntNode{val: 0}

	q.head.next = q.tail
	q.tail.prev = q.head

	return &q
}

// NewIntQueueFromValues создает новую очередь из значений
func NewIntQueueFromValues(values []int) *IntQueue {
	q := NewIntQueue()
	for _, v := range values {
		q.Enqueue(v)
	}
	return q
}

// Empty проверяет, пуста ли очередь
func (q *IntQueue) Empty() bool {
	return q.head.next == q.tail && q.tail.prev == q.head
}

// Enqueue добавляет элемент в конец очереди
func (q *IntQueue) Enqueue(val int) {
	node := &IntNode{val: val}
	node.next = q.tail
	node.prev = q.tail.prev
	q.tail.prev.next = node
	q.tail.prev = node
	q.Length++
}

// Peek возвращает первый элемент в очереди, но не удаляет его
func (q *IntQueue) Peek() int {
	return q.head.next.val
}

// Dequeue удаляет первый элемент в очереди и возвращает его значение
func (q *IntQueue) Dequeue() int {
	firstNode := q.head.next
	firstNode.next.prev = q.head
	q.head.next = firstNode.next
	q.Length--
	return firstNode.val
}

func countStudentsV3(students []int, sandwiches []int) int {
	/*
		METHOD: Queue
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	studentQueue, sandwichQueue := NewIntQueueFromValues(students), NewIntQueueFromValues(sandwiches)

	refusalStreak := 0
	for !studentQueue.Empty() && refusalStreak < studentQueue.Length {
		student := studentQueue.Dequeue()
		sandwich := sandwichQueue.Peek()

		if student != sandwich { // студент посмотрел на сендвич и не хочет его, идет в конец очереди
			studentQueue.Enqueue(student)
			refusalStreak++
		} else { // студент хочет сендвич, берет его из стека
			sandwichQueue.Dequeue()
			refusalStreak = 0
		}
	}

	return refusalStreak
}
