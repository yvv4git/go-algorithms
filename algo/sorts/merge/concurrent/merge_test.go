package concurrent

import (
	"fmt"
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
	}{
		{
			name: "CASE-1",
			args: args{
				data: []int{22, 8, 3, 31, 4, 2, 42, 1, 16, 6, 11, 25, 9, 8, 10, 12, 18, 14, 7, 15},
				r:    make(chan []int),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make(chan []int)

			go MergeSort(tt.args.data, result)

			r := <-result
			for _, v := range r {
				fmt.Println(v)
			}

			close(result)
		})
	}
}
