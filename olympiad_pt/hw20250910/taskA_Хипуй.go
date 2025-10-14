//go:build ignore

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MaxHeap struct {
	data []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		data: make([]int, 0),
	}
}

func (h *MaxHeap) Insert(k int) {
	h.data = append(h.data, k)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeap) Extract() int {
	if len(h.data) == 0 {
		return 0
	}

	max := h.data[0]
	last := len(h.data) - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]

	if len(h.data) > 0 {
		h.heapifyDown(0)
	}

	return max
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent] >= h.data[index] {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	size := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < size && h.data[left] > h.data[largest] {
			largest = left
		}
		if right < size && h.data[right] > h.data[largest] {
			largest = right
		}
		if largest == index {
			break
		}

		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	heap := NewMaxHeap()

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		switch parts[0] {
		case "0":
			// Insert operation
			num, _ := strconv.Atoi(parts[1])
			heap.Insert(num)
		case "1":
			// Extract operation
			result := heap.Extract()
			fmt.Println(result)
		}
	}
}
