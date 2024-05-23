package main

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Increasing sequence at the beginning",
			args: args{nums: []int{1, 2, 3, 1, 2, 3}},
			want: 3,
		},
		{
			name: "Test Case 2: Single element",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "Test Case 3: Empty array",
			args: args{nums: []int{}},
			want: 0,
		},
		{
			name: "Test Case 4: No increasing sequence",
			args: args{nums: []int{5, 4, 3, 2, 1}},
			want: 1,
		},
		{
			name: "Test Case 5: All elements are the same",
			args: args{nums: []int{1, 1, 1, 1, 1}},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCIS(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
