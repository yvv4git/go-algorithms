package main

import (
	"fmt"
)

// Структура для хранения элементов очереди с приоритетами
type PriorityQueue struct {
	elements []int
}

// NewPriorityQueue создает новую очередь с приоритетами
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

// Push добавляет элемент в очередь с приоритетами
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (pq *PriorityQueue) Push(value int) {
	pq.elements = append(pq.elements, value)
	pq.heapifyUp(len(pq.elements) - 1)
}

// Pop удаляет и возвращает элемент с наивысшим приоритетом из очереди
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (pq *PriorityQueue) Pop() (int, error) {
	if len(pq.elements) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	root := pq.elements[0]
	lastIndex := len(pq.elements) - 1
	pq.elements[0] = pq.elements[lastIndex]
	pq.elements = pq.elements[:lastIndex]
	pq.heapifyDown(0)
	return root, nil
}

// heapifyUp поддерживает свойство кучи при добавлении элемента
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (pq *PriorityQueue) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.elements[parentIndex] <= pq.elements[index] {
			break
		}
		pq.elements[parentIndex], pq.elements[index] = pq.elements[index], pq.elements[parentIndex]
		index = parentIndex
	}
}

// heapifyDown поддерживает свойство кучи при удалении элемента
// Временная сложность: O(log n)
// Пространственная сложность: O(1)
func (pq *PriorityQueue) heapifyDown(index int) {
	length := len(pq.elements)
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallest := index

		if leftChildIndex < length && pq.elements[leftChildIndex] < pq.elements[smallest] {
			smallest = leftChildIndex
		}
		if rightChildIndex < length && pq.elements[rightChildIndex] < pq.elements[smallest] {
			smallest = rightChildIndex
		}
		if smallest == index {
			break
		}
		pq.elements[smallest], pq.elements[index] = pq.elements[index], pq.elements[smallest]
		index = smallest
	}
}

// IsEmpty проверяет, пуста ли очередь
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.elements) == 0
}

// Peek возвращает элемент с наивысшим приоритетом без удаления из очереди
// Временная сложность: O(1)
// Пространственная сложность: O(1)
func (pq *PriorityQueue) Peek() (int, error) {
	if len(pq.elements) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return pq.elements[0], nil
}

func main() {
	pq := NewPriorityQueue()
	pq.Push(3)
	pq.Push(2)
	pq.Push(1)
	pq.Push(4)

	for !pq.IsEmpty() {
		val, _ := pq.Pop()
		fmt.Println(val)
	}
}
