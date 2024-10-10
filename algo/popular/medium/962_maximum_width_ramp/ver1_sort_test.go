package main

import (
	"testing"
)

func TestMaxWidthRamp(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Single element",
			nums: []int{1},
			want: 0,
		},
		{
			name: "Two elements",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "Large input",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want: 9,
		},
		{
			name: "No ramp",
			nums: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			want: 0,
		},
		{
			name: "Mixed input",
			nums: []int{6, 0, 8, 2, 1, 5},
			want: 4,
		},
		{
			name: "Duplicate values",
			nums: []int{1, 1, 1, 1, 1},
			want: 4,
		},
		{
			name: "Negative values",
			nums: []int{-1, -2, -3, -4, -5},
			want: 0,
		},
		{
			name: "Mixed positive and negative values",
			nums: []int{-1, 2, -3, 4, -5, 6},
			want: 5,
		},
		{
			name: "Large ramp",
			nums: []int{10, 0, 20, 0, 30, 0, 40, 0, 50},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthRamp(tt.nums); got != tt.want {
				t.Errorf("maxWidthRampV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
