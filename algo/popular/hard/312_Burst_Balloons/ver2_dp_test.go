package main

import "testing"

func Test_maxCoinsV2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 1, 5, 8},
			want: 167,
		},
		{
			name: "Example 2",
			nums: []int{1, 5},
			want: 10,
		},
		{
			name: "Single balloon",
			nums: []int{7},
			want: 7,
		},
		{
			name: "Two balloons",
			nums: []int{3, 8},
			want: 32,
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxCoinsV2(tt.nums); got != tt.want {
				t.Errorf("maxCoinsV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
