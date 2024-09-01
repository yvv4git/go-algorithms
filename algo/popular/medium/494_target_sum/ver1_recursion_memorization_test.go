package main

import "testing"

func Test_findTargetSumWaysV2(t *testing.T) {
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
			name: "Simple case with positive target",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 3,
			},
			want: 5,
		},
		{
			name: "Simple case with zero target",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: 0,
			},
			want: 0,
		},
		{
			name: "Case with negative target",
			args: args{
				nums:   []int{1, 1, 1, 1, 1},
				target: -3,
			},
			want: 5,
		},
		{
			name: "Case with single element array",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 1,
		},
		{
			name: "Case with no possible ways",
			args: args{
				nums:   []int{1, 2, 3},
				target: 7,
			},
			want: 0,
		},
		{
			name: "Case with all elements zero",
			args: args{
				nums:   []int{0, 0, 0, 0, 0},
				target: 0,
			},
			want: 32, // 2^5 ways to choose signs for zeros
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTargetSumWaysV2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findTargetSumWaysV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
