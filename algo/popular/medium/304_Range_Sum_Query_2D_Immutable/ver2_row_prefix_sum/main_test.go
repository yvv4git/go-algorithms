package main

import "testing"

func Test_SumRegionV2(t *testing.T) {
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
		{
			name: "Matrix 1 x 200",
			matrix: [][]int{
				{
					1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
					11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
					21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
					31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
					41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
					51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
					61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
					71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
					81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
					91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
					101, 102, 103, 104, 105, 106, 107, 108, 109, 110,
					111, 112, 113, 114, 115, 116, 117, 118, 119, 120,
					121, 122, 123, 124, 125, 126, 127, 128, 129, 130,
					131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
					141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
					151, 152, 153, 154, 155, 156, 157, 158, 159, 160,
					161, 162, 163, 164, 165, 166, 167, 168, 169, 170,
					171, 172, 173, 174, 175, 176, 177, 178, 179, 180,
					181, 182, 183, 184, 185, 186, 187, 188, 189, 190,
					191, 192, 193, 194, 195, 196, 197, 198, 199, 200,
				},
			},
			row1: 0, col1: 0, row2: 0, col2: 199,
			want: 20100,
		},
		{
			name: "Matrix 200 x 1",
			matrix: func() [][]int {
				matrix := make([][]int, 200)
				for i := range matrix {
					matrix[i] = []int{i + 1}
				}
				return matrix
			}(),
			row1: 0, col1: 0, row2: 199, col2: 0,
			want: 20100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := ConstructorV2(tt.matrix)
			if got := nm.SumRegionV2(tt.row1, tt.col1, tt.row2, tt.col2); got != tt.want {
				t.Errorf("SumRegionV2(%d, %d, %d, %d) = %v, want %v",
					tt.row1, tt.col1, tt.row2, tt.col2, got, tt.want)
			}
		})
	}
}
