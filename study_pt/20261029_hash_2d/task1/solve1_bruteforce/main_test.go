package main

import "testing"

func Test_countDistinctSets(t *testing.T) {
	tests := []struct {
		name string
		sets [][]int
		want int
	}{
		{
			name: "example 1",
			sets: [][]int{{1, 2, 3}, {3, 2, 1}, {11, 12, 13}, {2, 1, 3}},
			want: 2,
		},
		{
			name: "empty input",
			sets: [][]int{},
			want: 0,
		},
		{
			name: "single set",
			sets: [][]int{{5, 7, 9}},
			want: 1,
		},
		{
			name: "all sets distinct",
			sets: [][]int{{1}, {2}, {3}},
			want: 3,
		},
		{
			name: "single element sets with duplicates",
			sets: [][]int{{1}, {1}, {1}, {2}},
			want: 2,
		},
		{
			name: "same set in different order",
			sets: [][]int{{1, 2}, {2, 1}, {1, 2}, {2, 1}},
			want: 1,
		},
		{
			name: "different lengths",
			sets: [][]int{{1}, {1, 2}, {2, 1}, {1, 2, 3}},
			want: 3,
		},
		{
			name: "all identical values across sets",
			sets: [][]int{{0, 1, 2}, {2, 1, 0}, {1, 0, 2}},
			want: 1,
		},
		{
			name: "duplicate sets after sorting",
			sets: [][]int{{7, 8, 9}, {9, 8, 7}, {7, 9, 8}, {10, 11, 12}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDistinctSets(tt.sets); got != tt.want {
				t.Fatalf("countDistinctSets() = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_countDistinctSets_maxSize(t *testing.T) {
	const k = 1000000
	row := make([]int, k)
	for i := range row {
		row[i] = i
	}
	sets := [][]int{row}

	if got := countDistinctSets(sets); got != 1 {
		t.Fatalf("countDistinctSets() = %d, want 1", got)
	}
}
