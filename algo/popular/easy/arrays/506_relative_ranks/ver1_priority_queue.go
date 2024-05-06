package main

import (
	"container/heap"
	"fmt"
)

// Score содержит информацию о результате участника и его индекс в исходном массиве
type Score struct {
	Value int
	Index int
}

// PriorityQueue - приоритетная очередь, реализованная с помощью кучи
type PriorityQueue []*Score

// Len возвращает количество элементов в очереди
func (pq PriorityQueue) Len() int { return len(pq) }

// Less определяет, какой элемент имеет больший приоритет
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Value > pq[j].Value
}

// Swap меняет местами элементы с индексами i и j
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push добавляет элемент в очередь
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Score)
	*pq = append(*pq, item)
}

// Pop удаляет и возвращает элемент с наивысшим приоритетом из очереди
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // избегаем утечки памяти
	*pq = old[0 : n-1]
	return item
}

// findRelativeRanks находит относительные ранги участников и возвращает их в виде массива строк
func findRelativeRanks(score []int) []string {
	/*
		METHOD: Priority queue
		Использование кучи (Priority Queue) позволяет нам эффективно находить элемент с наивысшим приоритетом (наибольший результат) и удалять его из очереди.
		Это позволяет нам отсортировать результаты участников в порядке убывания, что позволяет нам легко определить их относительные ранги.

		TIME COMPLEXITY: O(n log n), где n - количество участников.
		Это связано с тем, что мы используем кучу для сортировки результатов и потом проходим по отсортированному массиву.

		SPACE COMPLEXITY: O(n), так как мы храним результаты в массиве и используем кучу для сортировки.
	*/
	pq := make(PriorityQueue, len(score))
	result := make([]string, len(score))

	// Заполняем очередь и сохраняем исходные индексы
	for i, num := range score {
		pq[i] = &Score{Value: num, Index: i}
	}
	heap.Init(&pq)

	// Находим относительные ранги
	for i := 0; pq.Len() > 0; i++ {
		score := heap.Pop(&pq).(*Score)
		switch i {
		case 0:
			result[score.Index] = "Gold Medal"
		case 1:
			result[score.Index] = "Silver Medal"
		case 2:
			result[score.Index] = "Bronze Medal"
		default:
			result[score.Index] = fmt.Sprintf("%d", i+1)
		}
	}

	return result
}

func main() {
	nums := []int{5, 4, 3, 2, 1}
	ranks := findRelativeRanks(nums)
	fmt.Println(ranks)
}
