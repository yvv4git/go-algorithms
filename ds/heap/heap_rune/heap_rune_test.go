package heap_rune

import (
	"container/heap"
	"testing"
)

func TestHeapRune(t *testing.T) {
	a := []rune("Vladimir")
	h := runeSlice(a)
	heap.Init(&h)
	t.Logf("a=%v a_s=%v", a, string(a)) // a=[114 108 109 100 105 97 105 86] a_s=rlmdiaiV
}
