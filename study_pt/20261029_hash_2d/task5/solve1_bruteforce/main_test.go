package main

import "testing"

func Test_countNiceRectangles(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		grid []string
		want int
	}{
		{
			name: "example 1",
			n:    2,
			m:    3,
			grid: []string{"abc", "acb"},
			want: 8,
		},
		{
			name: "single cell",
			n:    1,
			m:    1,
			grid: []string{"a"},
			want: 1,
		},
		{
			name: "two different cells in a row",
			n:    1,
			m:    2,
			grid: []string{"ab"},
			want: 2,
		},
		{
			name: "two equal cells in a row",
			n:    1,
			m:    2,
			grid: []string{"aa"},
			want: 3,
		},
		{
			name: "palindrome row",
			n:    1,
			m:    3,
			grid: []string{"aba"},
			want: 4,
		},
		{
			name: "all cells equal 2x2",
			n:    2,
			m:    2,
			grid: []string{"aa", "aa"},
			want: 9,
		},
		{
			name: "symmetric 2x2",
			n:    2,
			m:    2,
			grid: []string{"ab", "ba"},
			want: 5,
		},
		{
			name: "all cells different 2x2",
			n:    2,
			m:    2,
			grid: []string{"ab", "cd"},
			want: 4,
		},
		{
			name: "all cells equal 3x3",
			n:    3,
			m:    3,
			grid: []string{"aaa", "aaa", "aaa"},
			want: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNiceRectangles(tt.n, tt.m, tt.grid); got != tt.want {
				t.Fatalf("countNiceRectangles() = %d, want %d", got, tt.want)
			}
		})
	}
}
