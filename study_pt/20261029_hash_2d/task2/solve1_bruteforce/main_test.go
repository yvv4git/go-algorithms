package main

import "testing"

func Test_countIsomorphicSubtreePairs(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  [][2]int
		values []int
		want   int
	}{
		{
			name:   "example 1",
			n:      4,
			edges:  [][2]int{{1, 2}, {1, 3}, {3, 4}},
			values: []int{1, 2, 3, 2},
			want:   1,
		},
		{
			name:   "example 2",
			n:      6,
			edges:  [][2]int{{1, 2}, {1, 3}, {1, 4}, {3, 5}, {4, 6}},
			values: []int{1, 2, 3, 4, 2, 2},
			want:   3,
		},
		{
			name:   "two equal leaf children",
			n:      3,
			edges:  [][2]int{{1, 2}, {1, 3}},
			values: []int{7, 7, 7},
			want:   1,
		},
		{
			name:   "two different leaves",
			n:      2,
			edges:  [][2]int{{1, 2}},
			values: []int{7, 8},
			want:   0,
		},
		{
			name:   "chain with equal values",
			n:      4,
			edges:  [][2]int{{1, 2}, {2, 3}, {3, 4}},
			values: []int{5, 5, 5, 5},
			want:   0,
		},
		{
			name:   "star with equal leaves",
			n:      4,
			edges:  [][2]int{{1, 2}, {1, 3}, {1, 4}},
			values: []int{1, 2, 2, 2},
			want:   3,
		},
		{
			name:   "star with two equal leaves",
			n:      4,
			edges:  [][2]int{{1, 2}, {1, 3}, {1, 4}},
			values: []int{1, 2, 2, 3},
			want:   1,
		},
		{
			name:   "all values distinct",
			n:      4,
			edges:  [][2]int{{1, 2}, {1, 3}, {1, 4}},
			values: []int{1, 2, 3, 4},
			want:   0,
		},
		{
			name:   "two equal inner subtrees and two equal leaves",
			n:      5,
			edges:  [][2]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}},
			values: []int{1, 2, 2, 3, 3},
			want:   2,
		},
		{
			name:   "big star with equal leaves",
			n:      6,
			edges:  [][2]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}},
			values: []int{1, 2, 2, 2, 2, 2},
			want:   10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countIsomorphicSubtreePairs(tt.n, tt.edges, tt.values); got != tt.want {
				t.Fatalf("countIsomorphicSubtreePairs() = %d, want %d", got, tt.want)
			}
		})
	}
}
