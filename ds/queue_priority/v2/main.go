package main

import (
	"container/heap"
	"fmt"
)

func main() {
	pq := &TaskPQ{
		{3, "Vacuum"},
		{2, "Feed cat"},
		{4, "Arrange play date"},
		{1, "Pack for trip"}}

	// Инициализируем
	heap.Init(pq)

	// Положить в очередь
	heap.Push(pq, Task{2, "Vladimir"})

	for pq.Len() != 0 {
		// Вытаскиваем из очереди
		fmt.Println(heap.Pop(pq))
	}
	/*
		{1 Pack for trip}
		{2 Vladimir}
		{2 Feed cat}
		{3 Vacuum}
		{4 Arrange play date}
	*/
}
