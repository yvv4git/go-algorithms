package squares_sorted_array

import (
	"reflect"
	"testing"
)

func Test_sortedSquaresV1(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{-4, -1, 0, 3, 10},
			},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{-7, -3, 2, 3, 11},
			},
			want: []int{4, 9, 9, 49, 121},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquaresV1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquaresV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
