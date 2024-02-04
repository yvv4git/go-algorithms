package max_number_events_attended

import (
	"container/heap"
	"sort"
)

func maxEventsV1(events [][]int) int {
	/*
		METHOD: Heap
		TIME COMPLEXITY: ???
		Space complexity: ???
	*/

	// Convert the events into Event struct for easy sorting
	var eventsList []Event
	for _, event := range events {
		eventsList = append(eventsList, Event{event[0], event[1]})
	}

	// Sort events according to start time
	sortEventList(eventsList)

	totalDays := 0
	for _, event := range eventsList {
		if event.endTime > totalDays {
			totalDays = event.endTime
		}
	}

	minHeap := &EventHeap{}
	heap.Init(minHeap)
	day, cnt, eventID := 1, 0, 0

	for day <= totalDays {
		// If no events are available to attend today, let time flies to the next available event.
		if eventID < len(eventsList) && minHeap.Len() == 0 {
			day = eventsList[eventID].startTime
		}

		// All events starting from today are newly available. Add them to the heap.
		for eventID < len(eventsList) && eventsList[eventID].startTime <= day {
			heap.Push(minHeap, eventsList[eventID].endTime)
			eventID++
		}

		// If the event at heap top already ended, then discard it.
		for minHeap.Len() > 0 && (*minHeap).items[0] < day {
			heap.Pop(minHeap)
		}

		// Attend the event that will end the earliest
		if minHeap.Len() > 0 {
			heap.Pop(minHeap)
			cnt++
		} else if eventID >= len(eventsList) {
			break // No more events to attend, so stop early to save time.
		}
		day++
	}

	return cnt
}

func sortEventList(list []Event) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].startTime < list[j].startTime
	})
}

type Event struct {
	startTime, endTime int
}

type EventHeap struct {
	items []int
}

func (h EventHeap) Len() int {
	return len(h.items)
}

func (h EventHeap) Less(i, j int) bool {
	return h.items[i] < h.items[j]
}

func (h EventHeap) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *EventHeap) Push(x interface{}) {
	h.items = append(h.items, x.(int))
}

func (h *EventHeap) Pop() interface{} {
	old := h.items
	n := len(old)
	x := old[n-1]
	h.items = old[0 : n-1]
	return x
}
