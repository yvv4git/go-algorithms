package top_k_frequently_elements

import (
	"reflect"
	"testing"
)

func Test_topKFrequentV2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{1, 1, 1, 1, 2, 3, 3, 3, 4, 4, 1},
				k:    3,
			},
			want: []int{1, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequentV2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequentV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
