package max_number_events_attended

import (
	"container/heap"
	"sort"
)

func maxEventsV2(events [][]int) int {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	pq := new(priorityQueue)
	eventsIndex, out := 0, 0

	for currentDay := 1; currentDay <= 100000; currentDay++ {
		for !pq.IsEmpty() && pq.Peek().(int) < currentDay {
			heap.Pop(pq)
		}

		for eventsIndex < len(events) && events[eventsIndex][0] == currentDay {
			heap.Push(pq, events[eventsIndex][1])
			eventsIndex++
		}

		if !pq.IsEmpty() {
			heap.Pop(pq)
			out++
		}
	}

	return out
}

type priorityQueue []int

func (pq priorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

// Less - use min priority heap.
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(num interface{}) {
	*pq = append(*pq, num.(int))
}

func (pq *priorityQueue) Pop() interface{} {
	l := len(*pq)
	num := (*pq)[l-1]
	*pq = (*pq)[:l-1]

	return num
}

func (pq priorityQueue) Peek() interface{} {
	return pq[0]
}
