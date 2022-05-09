package heapbinary

// BinaryHeap
type BinaryHeap struct {
	heap []int
}

// Push - implemetation push method
func (bh *BinaryHeap) Push(key int) {
	bh.heap = append(bh.heap, key)
	bh.bubbleUp(len(bh.heap) - 1)
}

func (bh *BinaryHeap) bubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2

		if bh.heap[parentIndex] < bh.heap[index] {
			return
		}

		bh.heap[parentIndex], bh.heap[index] = bh.heap[index], bh.heap[parentIndex]
		index = parentIndex
	}
}

// Pop - implementation pop method
func (bh *BinaryHeap) Pop() int {
	popped := bh.heap[0]
	heapSize := len(bh.heap)

	if heapSize > 1 {
		bh.heap[0] = bh.heap[len(bh.heap)-1]
	}

	bh.heap = bh.heap[:len(bh.heap)-1]
	bh.bubbleDown(0)
	return popped
}

func (bh *BinaryHeap) bubbleDown(index int) {
	for 2*index+1 < len(bh.heap) {
		minChildIndex := bh.minChildIndex(index)

		if bh.heap[minChildIndex] > bh.heap[index] {
			return
		}

		bh.heap[minChildIndex], bh.heap[index] = bh.heap[index], bh.heap[minChildIndex]
		index = minChildIndex
	}
}

func (bh *BinaryHeap) minChildIndex(index int) int {
	if 2*index+2 >= len(bh.heap) {
		return 2*index + 1
	}

	if bh.heap[2*index+2] < bh.heap[2*index+1] {
		return 2*index + 2
	}

	return 2*index + 1
}
