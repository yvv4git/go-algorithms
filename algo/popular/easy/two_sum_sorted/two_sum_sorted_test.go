package main

import "testing"

func TestTwoSumSorted(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CASE-1",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 12,
			},
			want: true,
		},
		{
			name: "CASE-2",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				target: 1,
			},
			want: false,
		},
		{
			name: "CASE-3",
			args: args{
				nums:   []int{},
				target: 100500,
			},
			want: false,
		},
		{
			name: "CASE-4",
			args: args{
				nums:   []int{-5, -4, 0, 1, 3, 6, 7, 8, 9},
				target: 7,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumSorted(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("TwoSumSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
