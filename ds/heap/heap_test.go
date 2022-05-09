package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)

	for h.Len() > 0 {
		fmt.Printf("%d", heap.Pop(h))
	}
}
