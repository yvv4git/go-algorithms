package _5_3sum

import (
	"reflect"
	"testing"
)

func Test_threeSumV4(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{0, 0, 0},
			},
			want: [][]int{{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumV4(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSumV4() = %v, want %v", got, tt.want)
			}
		})
	}
}
