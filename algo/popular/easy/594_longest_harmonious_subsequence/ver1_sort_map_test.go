package main

import "testing"

func Test_findLHS(t *testing.T) {
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
			args: args{nums: []int{1, 3, 2, 2, 5, 2, 3, 7}},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{nums: []int{1, 2, 2, 1}},
			want: 4,
		},
		{
			name: "Example 5",
			args: args{nums: []int{1, 3, 5, 7, 9}},
			want: 0,
		},
		{
			name: "Example 6",
			args: args{nums: []int{1, 1, 2, 2, 2, 2, 3, 3, 3, 3}},
			want: 8,
		},
		{
			name: "Example 7",
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			name: "Example 8",
			args: args{nums: []int{1, 1, 1, 2, 2, 2}},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLHS(tt.args.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
