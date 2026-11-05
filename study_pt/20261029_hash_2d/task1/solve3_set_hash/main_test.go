package main

import "testing"

func Test_countDistinctSetsV3(t *testing.T) {
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
			name: "large values",
			sets: [][]int{{1000000000, 0}, {0, 1000000000}, {999999999, 1}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDistinctSetsV3(tt.sets); got != tt.want {
				t.Fatalf("countDistinctSetsV3() = %d, want %d", got, tt.want)
			}
		})
	}
}

func Test_countDistinctSetsV3_maxSize(t *testing.T) {
	const n = 100000
	const k = 10
	sets := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, k)
		for j := 0; j < k; j++ {
			row[j] = i*k + j
		}
		sets[i] = row
	}

	if got := countDistinctSetsV3(sets); got != n {
		t.Fatalf("countDistinctSetsV3() = %d, want %d", got, n)
	}
}
