package _470_shuffle_the_array

import (
	"reflect"
	"testing"
)

func Test_shuffleV2(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}

	tests := []struct {
		name string
		args args
		want []int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{2, 5, 1, 3, 4, 7},
				n:    3,
			},
			want: []int{2, 3, 5, 4, 1, 7},
			desc: `Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7 then the answer is [2,3,5,4,1,7]`,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
				n:    4,
			},
			want: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{1, 1, 2, 2},
				n:    2,
			},
			want: []int{1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffleV2(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shuffleV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
