package main

import "testing"

func Test_wiggleSortV3(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 5, 1, 1, 6, 4}},
			want: []int{1, 6, 1, 5, 1, 4},
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 3, 2, 2, 3, 1}},
			want: []int{2, 3, 1, 3, 1, 2},
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 1, 2, 1, 2, 2, 1}},
			want: []int{1, 2, 1, 2, 1, 2, 1},
		},
		{
			name: "Example 4",
			args: args{nums: []int{4, 5, 5, 6}},
			want: []int{5, 6, 4, 5},
		},
		{
			name: "Example 5",
			args: args{nums: []int{1}},
			want: []int{1},
		},
		{
			name: "Example 6",
			args: args{nums: []int{}},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wiggleSortV3(tt.args.nums)
		})
	}
}
