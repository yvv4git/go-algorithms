package main

import "testing"

func Test_SumRegionV3(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		row1   int
		col1   int
		row2   int
		col2   int
		want   int
	}{
		{
			name: "Example 1: sumRegion(2, 1, 4, 3)",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			row1: 2, col1: 1, row2: 4, col2: 3,
			want: 8,
		},
		{
			name: "Example 1: sumRegion(1, 1, 2, 2)",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			row1: 1, col1: 1, row2: 2, col2: 2,
			want: 11,
		},
		{
			name: "Example 1: sumRegion(1, 2, 2, 4)",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			row1: 1, col1: 2, row2: 2, col2: 4,
			want: 12,
		},
		{
			name: "Example 2: sumRegion(0, 0, 1, 1)",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			row1: 0, col1: 0, row2: 1, col2: 1,
			want: 12,
		},
		{
			name: "Example 2: sumRegion(1, 1, 2, 2)",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			row1: 1, col1: 1, row2: 2, col2: 2,
			want: 28,
		},
		{
			name:   "Single element matrix",
			matrix: [][]int{{7}},
			row1:   0, col1: 0, row2: 0, col2: 0,
			want: 7,
		},
		{
			name:   "One row",
			matrix: [][]int{{1, 2, 3, 4}},
			row1:   0, col1: 1, row2: 0, col2: 3,
			want: 9,
		},
		{
			name:   "One column",
			matrix: [][]int{{1}, {2}, {3}},
			row1:   0, col1: 0, row2: 2, col2: 0,
			want: 6,
		},
		{
			name: "Whole matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			row1: 0, col1: 0, row2: 1, col2: 1,
			want: 10,
		},
		{
			name: "All zeros",
			matrix: [][]int{
				{0, 0},
				{0, 0},
			},
			row1: 0, col1: 0, row2: 1, col2: 1,
			want: 0,
		},
		{
			name: "Negative values",
			matrix: [][]int{
				{-1, -2, -3},
				{-4, -5, -6},
			},
			row1: 0, col1: 0, row2: 1, col2: 1,
			want: -12,
		},
		{
			name: "Duplicate values",
			matrix: [][]int{
				{5, 5, 5},
				{5, 5, 5},
			},
			row1: 0, col1: 0, row2: 1, col2: 2,
			want: 30,
		},
		{
			name: "Max border values",
			matrix: [][]int{
				{-10000, -10000},
				{-10000, -10000},
			},
			row1: 0, col1: 0, row2: 1, col2: 1,
			want: -40000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := ConstructorV3(tt.matrix)
			if got := nm.SumRegionV3(tt.row1, tt.col1, tt.row2, tt.col2); got != tt.want {
				t.Errorf("SumRegionV3(%d, %d, %d, %d) = %v, want %v",
					tt.row1, tt.col1, tt.row2, tt.col2, got, tt.want)
			}
		})
	}
}

// Test_SumRegionV3_MultipleQueries проверяет, что один построенный объект
// корректно отвечает на несколько запросов подряд: таблица префиксов
// строится один раз в конструкторе и дальше только читается.
func Test_SumRegionV3_MultipleQueries(t *testing.T) {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}

	nm := ConstructorV3(matrix)

	queries := []struct {
		name       string
		row1, col1 int
		row2, col2 int
		want       int
	}{
		{name: "Query 1", row1: 2, col1: 1, row2: 4, col2: 3, want: 8},
		{name: "Query 2", row1: 1, col1: 1, row2: 2, col2: 2, want: 11},
		{name: "Query 3", row1: 1, col1: 2, row2: 2, col2: 4, want: 12},
		{name: "Query 4: whole matrix", row1: 0, col1: 0, row2: 4, col2: 4, want: 58},
		{name: "Query 5: repeated Query 1", row1: 2, col1: 1, row2: 4, col2: 3, want: 8},
	}

	for _, q := range queries {
		t.Run(q.name, func(t *testing.T) {
			if got := nm.SumRegionV3(q.row1, q.col1, q.row2, q.col2); got != q.want {
				t.Errorf("SumRegionV3(%d, %d, %d, %d) = %v, want %v",
					q.row1, q.col1, q.row2, q.col2, got, q.want)
			}
		})
	}
}
