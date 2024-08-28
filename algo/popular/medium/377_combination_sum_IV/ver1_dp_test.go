package main

import "testing"

func Test_combinationSum4(t *testing.T) {
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
			name: "Example 1",
			args: args{
				nums:   []int{1, 2, 3},
				target: 4,
			},
			want: 7,
		},
		{
			name: "Example 2",
			args: args{
				nums:   []int{9},
				target: 3,
			},
			want: 0,
		},
		{
			name: "Example 3",
			args: args{
				nums:   []int{1, 2, 3},
				target: 32,
			},
			want: 181997601,
		},
		{
			name: "Example 4",
			args: args{
				nums:   []int{2, 3, 5},
				target: 8,
			},
			want: 6,
		},
		{
			name: "Example 5",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum4(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
			}
		})
	}
}
