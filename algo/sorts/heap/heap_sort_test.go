package heap

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		ar []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				ar: []int{2, 1},
			},
			want: []int{1, 2},
		},
		{
			name: "CASE-2",
			args: args{
				ar: []int{5, 1, 4, 2, 7},
			},
			want: []int{1, 2, 4, 5, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.ar)
			if !reflect.DeepEqual(tt.want, tt.args.ar) {
				t.Errorf("We expect value %v, but we got %v", tt.want, tt.args.ar)
			}
		})
	}
}
