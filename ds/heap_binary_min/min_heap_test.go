package heapbinarymin

import (
	"reflect"
	"testing"
)

func TestMinHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *MinHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MinHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		h    *MinHeap
		args args
	}{
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
		})
	}
}
