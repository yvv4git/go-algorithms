package top_k_frequently_elements

import (
	"container/heap"
)

func topKFrequentV4(nums []int, k int) []int {
	map1 := make(map[int]int) //value to frequency map
	for _, num := range nums {
		map1[num] = map1[num] + 1
	}

	h := &MinHeap{}
	heap.Init(h)

	//building heap with top k frequenct elements
	for key, value := range map1 {
		n := node{
			num:       key,
			frequency: value,
		}
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	// fmt.Println(h)
	//putting the elements in response array
	res := make([]int, 0)
	for h.Len() > 0 {
		result := heap.Pop(h)
		//fmt.Printf("%T",result.(node).num)

		res = append(res, result.(node).num)
	}

	return res
}

type node struct {
	num       int
	frequency int
}

type MinHeap []node

func (h MinHeap) Less(i, j int) bool {
	return h[i].frequency < h[j].frequency
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(node))
}

func (h *MinHeap) Pop() any {
	old := *h
	length := h.Len()
	x := old[length-1]
	*h = old[0 : length-1]
	return x
}
