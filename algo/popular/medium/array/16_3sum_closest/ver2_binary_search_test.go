package main

import "testing"

func Test_threeSumClosestV2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Exact match",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 2,
			},
			want: 2,
		},
		{
			name: "Test Case 2: Closest sum is the sum of two smallest numbers",
			args: args{
				nums:   []int{1, 1, -1, -1, 3},
				target: -1,
			},
			want: -1,
		},
		{
			name: "Test Case 3: Closest sum is in the middle",
			args: args{
				nums:   []int{-5, -2, 10, 20, -30, -40},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Test Case 4: Array with duplicate numbers",
			args: args{
				nums:   []int{1, 1, -1, -1, 3},
				target: -1,
			},
			want: -1,
		},
		{
			name: "Test Case 5: Array with all zeros",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosestV2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("threeSumClosestV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
