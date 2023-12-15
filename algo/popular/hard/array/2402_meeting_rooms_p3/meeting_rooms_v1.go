package _402_meeting_rooms_p3

import (
	"container/heap"
	"math"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	freeRooms := new(MinFreeHeap)
	busyRooms := new(MinBusyHeap)

	count := map[int]int{}
	for i := 0; i < n; i++ {
		count[i] = 0
	}

	for i := 0; i < n; i++ {
		rm := Room{
			Num:     i,
			EndTime: -1,
		}
		heap.Push(freeRooms, rm)
	}

	// Sort the order of the input, so that the meetings go one by one in the time order.
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	meetingSize := len(meetings)

	for i := 0; i < meetingSize; i++ {
		start := meetings[i][0]
		end := meetings[i][1]

		for busyRooms.Len() > 0 && busyRooms.MinFreeHeap[0].EndTime <= start {
			rm := heap.Pop(busyRooms).(Room)
			heap.Push(freeRooms, rm)
		}

		if freeRooms.Len() > 0 {
			rm := heap.Pop(freeRooms).(Room)
			count[rm.Num] += 1
			rm.EndTime = end
			// This Room is available, so the end time will be on time.
			heap.Push(busyRooms, rm)
		} else {
			rm := heap.Pop(busyRooms).(Room)
			count[rm.Num] += 1
			rm.EndTime += end - start
			heap.Push(busyRooms, rm)
		}
	}

	result := 0
	max := math.MinInt32
	for i := 0; i < n; i++ {
		if count[i] > max {
			max = count[i]
			result = i
		}
	}

	return result
}

type Room struct {
	Num     int
	EndTime int
}

type MinFreeHeap []Room

func (h *MinFreeHeap) Less(i, j int) bool {
	return (*h)[i].Num < (*h)[j].Num
}
func (h *MinFreeHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinFreeHeap) Len() int {
	return len(*h)
}

func (h *MinFreeHeap) Push(x interface{}) {
	*h = append(*h, x.(Room))
}

func (h *MinFreeHeap) Pop() interface{} {
	size := h.Len()
	first := (*h)[size-1]
	(*h) = (*h)[:size-1]

	return first
}

type MinBusyHeap struct {
	MinFreeHeap
}

func (h *MinBusyHeap) Less(i, j int) bool {
	res := h.MinFreeHeap[i].EndTime == h.MinFreeHeap[j].EndTime
	if res {
		return h.MinFreeHeap[i].Num < h.MinFreeHeap[j].Num
	} else {
		return h.MinFreeHeap[i].EndTime < h.MinFreeHeap[j].EndTime
	}
}
