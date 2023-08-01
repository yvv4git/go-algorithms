package priority_queue_on_heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	/*
		Priority Queue может быть реализована с помощью разных структур данных:
		- массив
		- связанный список
		- heap

		Самая эффективная для priority queue структура данных - heap. Так как обе операции(добавление и извлечение элемента)
		можно производить за O(log(n)) время. А вот с массивом и связанным списком эти операции могут требовать O(n).

		Основные функции:
		- добавить элемент в кучу
		- извлечь элемент из кучи
	*/
	items := map[int]int{
		1: 10, 2: 20, 3: 5,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))

	// Добавляем безмысленный элемент с 0 приоритетом.
	pq.Push(&Item{value: 0, priority: 0})

	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	fmt.Println(heap.Pop(&pq).(*Item))
	fmt.Println(heap.Pop(&pq).(*Item))
	fmt.Println(heap.Pop(&pq).(*Item))
	fmt.Println(heap.Pop(&pq).(*Item))
}
