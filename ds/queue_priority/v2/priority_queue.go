package main

type Task struct {
	priority int
	name     string
}

type TaskPQ []Task

// Len - возвращает длину очереди
func (q TaskPQ) Len() int {
	return len(q)
}

// Less - сравниваем значения двух элементов
func (q TaskPQ) Less(i, j int) bool {
	return q[i].priority < q[j].priority
}

// Swap - меняем местами элементы
func (q TaskPQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

// Push - помещаем в очередь
func (q *TaskPQ) Push(x interface{}) {
	*q = append(*q, x.(Task))
}

// Pop - достаем из очереди
func (q *TaskPQ) Pop() (popped interface{}) {
	popped = (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return
}
