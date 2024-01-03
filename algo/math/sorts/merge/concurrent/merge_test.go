package concurrent

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		data []int
		r    chan []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				data: []int{22, 8, 3, 31, 4, 2, 42, 1, 16, 6, 11, 25, 9, 8, 10, 12, 18, 14, 7, 15},
				r:    make(chan []int),
			},
			want: []int{1, 2, 3, 4, 6, 7, 8, 8, 9, 10, 11, 12, 14, 15, 16, 18, 22, 25, 31, 42},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chRes := make(chan []int)

			go MergeSort(tt.args.data, chRes)

			var result []int
			r := <-chRes
			for _, v := range r {
				result = append(result, v)
			}

			close(chRes)

			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("Sort() = %v, want %v", result, tt.want)
			}
		})
	}
}
