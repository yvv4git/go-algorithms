package main

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{10, 9, 2, 5, 3, 7, 101, 18}},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{nums: []int{0, 1, 0, 3, 2, 3}},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{nums: []int{7, 7, 7, 7, 7, 7, 7}},
			want: 1,
		},
		{
			name: "Empty Array",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "Single Element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "Increasing Sequence",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "Decreasing Sequence",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 1,
		},
		{
			name: "Mixed Sequence",
			args: args{nums: []int{10, 22, 9, 33, 21, 50, 41, 60, 80}},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
