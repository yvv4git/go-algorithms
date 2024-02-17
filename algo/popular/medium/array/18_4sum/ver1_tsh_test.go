package main

import (
	"reflect"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name: "Test Case 2",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			want: [][]int{{2, 2, 2, 2}},
		},
		{
			name: "Test Case 3",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2, 3, -3},
				target: 0,
			},
			want: [][]int{{-3, -2, 2, 3}, {-3, -1, 1, 3}, {-3, 0, 0, 3}, {-3, 0, 1, 2}, {-2, -1, 0, 3}, {-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name: "Test Case 4",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 10,
			},
			want: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "Test Case 5",
			args: args{
				nums:   []int{0, 0, 0, 0},
				target: 0,
			},
			want: [][]int{{0, 0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
