package heapbinary

import (
	"testing"
)

func TestHeap(t *testing.T) {
	testCasess := []struct {
		name     string
		elements []int
	}{
		{
			name:     "CASE-1",
			elements: []int{5, 3, 2, 4, 1},
		},
	}

	for _, tc := range testCasess {
		t.Run(tc.name, func(t *testing.T) {
			h := BinaryHeap{}
			for _, num := range tc.elements {
				h.Push(num)
			}

			t.Logf("%v", h.heap)
		})
	}
}

func TestBinaryHeap_Push(t *testing.T) {
	type fields struct {
		heap []int
	}
	type args struct {
		key int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bh := &BinaryHeap{
				heap: tt.fields.heap,
			}
			bh.Push(tt.args.key)
		})
	}
}
