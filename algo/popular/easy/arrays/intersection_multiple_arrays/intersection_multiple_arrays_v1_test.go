package intersection_multiple_arrays

import (
	"reflect"
	"testing"
)

func Test_intersectionV1(t *testing.T) {
	type args struct {
		nums [][]int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "CASE-1",
			args: args{[][]int{
				{3, 1, 2, 4, 5},
				{1, 2, 3, 4},
				{3, 4, 5, 6},
			}},
			want: []int{3, 4},
		},
		{
			name: "CASE-2",
			args: args{[][]int{
				{1, 2, 3},
				{4, 5, 6},
			}},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersectionV1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersectionV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
