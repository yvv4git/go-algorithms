package main

import (
	"testing"
)

func TestDistributeCandiesV2(t *testing.T) {
	tests := []struct {
		name      string
		candyType []int
		want      int
	}{
		{
			name:      "example 1",
			candyType: []int{1, 1, 2, 2, 3, 3},
			want:      3,
		},
		{
			name:      "example 2",
			candyType: []int{1, 1, 2, 3},
			want:      2,
		},
		{
			name:      "all same",
			candyType: []int{1, 1, 1, 1},
			want:      1,
		},
		{
			name:      "all different",
			candyType: []int{1, 2, 3, 4},
			want:      2,
		},
		{
			name:      "empty",
			candyType: []int{},
			want:      0,
		},
		{
			name:      "large input",
			candyType: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			want:      5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := distributeCandiesV2(tt.candyType); got != tt.want {
				t.Errorf("distributeCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
