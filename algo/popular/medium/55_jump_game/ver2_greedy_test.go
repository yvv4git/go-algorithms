package main

import "testing"

func Test_canJumpV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{nums: []int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "Example 2",
			args: args{nums: []int{3, 2, 1, 0, 4}},
			want: false,
		},
		{
			name: "Single element",
			args: args{nums: []int{0}},
			want: true,
		},
		{
			name: "All zeros",
			args: args{nums: []int{0, 0, 0, 0}},
			want: false,
		},
		{
			name: "Jump to end",
			args: args{nums: []int{1, 2, 3}},
			want: true,
		},
		{
			name: "Large jumps",
			args: args{nums: []int{5, 0, 0, 0, 0, 0}},
			want: true,
		},
		{
			name: "No way to reach end",
			args: args{nums: []int{1, 1, 0, 1}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJumpV2(tt.args.nums); got != tt.want {
				t.Errorf("canJumpV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
