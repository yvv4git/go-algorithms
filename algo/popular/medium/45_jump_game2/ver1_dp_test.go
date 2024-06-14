package main

import "testing"

func Test_jump(t *testing.T) {
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
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{nums: []int{2, 3, 0, 1, 4}},
			want: 2,
		},
		{
			name: "Single element",
			args: args{nums: []int{1}},
			want: 0,
		},
		{
			name: "All elements are 1",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 3,
		},
		{
			name: "Large jumps",
			args: args{nums: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}},
			want: 3,
		},
		{
			name: "No jumps needed",
			args: args{nums: []int{0}},
			want: 0,
		},
		{
			name: "Increasing jumps",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump(tt.args.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
