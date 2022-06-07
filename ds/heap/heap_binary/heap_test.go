package heapbinary

import (
	"reflect"
	"testing"
)

func TestBinaryHeap_Push(t *testing.T) {
	type args struct {
		elements []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				elements: []int{5, 3, 2, 4, 1},
			},
			want: []int{1, 2, 3, 5, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bh := BinaryHeap{}
			for _, val := range tt.args.elements {
				bh.Push(val)
			}

			if !reflect.DeepEqual(bh.heap, tt.want) {
				t.Errorf("BinaryHeap.Push() want %v, but created %v", tt.want, bh.heap)
			}
		})
	}
}

func TestBinaryHeap_Pop(t *testing.T) {
	type fields struct {
		keys []int
	}

	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "CASE-1",
			fields: fields{
				keys: []int{1, 2, 3, 5, 4},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bh := &BinaryHeap{
				heap: tt.fields.keys,
			}
			if got := bh.Pop(); got != tt.want {
				t.Errorf("BinaryHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
