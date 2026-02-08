package main

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		N    int
		M    int
		grid [][]int
	}

	tests := []struct {
		name        string
		args        args
		wantMaxSum  int
		wantPathLen int // длина пути должна быть N+M-2
	}{
		{
			name: "Example from README",
			args: args{
				N: 5,
				M: 5,
				grid: [][]int{
					{9, 9, 9, 9, 9},
					{3, 0, 0, 0, 0},
					{9, 9, 9, 9, 9},
					{6, 6, 6, 6, 8},
					{9, 9, 9, 9, 9},
				},
			},
			wantMaxSum:  74,
			wantPathLen: 8, // 5+5-2 = 8
		},
		{
			name: "1x1 grid",
			args: args{
				N:    1,
				M:    1,
				grid: [][]int{{5}},
			},
			wantMaxSum:  5,
			wantPathLen: 0, // 1+1-2 = 0
		},
		{
			name: "1x2 grid",
			args: args{
				N:    1,
				M:    2,
				grid: [][]int{{1, 2}},
			},
			wantMaxSum:  3,
			wantPathLen: 1, // 1+2-2 = 1
		},
		{
			name: "2x1 grid",
			args: args{
				N:    2,
				M:    1,
				grid: [][]int{{1}, {2}},
			},
			wantMaxSum:  3,
			wantPathLen: 1, // 2+1-2 = 1
		},
		{
			name: "2x2 grid",
			args: args{
				N:    2,
				M:    2,
				grid: [][]int{{1, 2}, {3, 4}},
			},
			wantMaxSum:  8, // 1->3->4 = 1+3+4 = 8
			wantPathLen: 2, // 2+2-2 = 2
		},
		{
			name: "All zeros",
			args: args{
				N: 3,
				M: 3,
				grid: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			wantMaxSum:  0,
			wantPathLen: 4, // 3+3-2 = 4
		},
		{
			name: "All ones",
			args: args{
				N: 3,
				M: 3,
				grid: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			wantMaxSum:  5, // 3+3-2 = 5 ячеек
			wantPathLen: 4,
		},
		{
			name: "Prefer right movement",
			args: args{
				N: 2,
				M: 3,
				grid: [][]int{
					{1, 100, 100},
					{1, 1, 1},
				},
			},
			wantMaxSum:  202, // 1->100->100->1 = 202 (максимум по правой стороне)
			wantPathLen: 3,   // 2+3-2 = 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMaxSum, gotPath := solve(tt.args.N, tt.args.M, tt.args.grid)
			if gotMaxSum != tt.wantMaxSum {
				t.Errorf("solve() maxSum = %v, want %v", gotMaxSum, tt.wantMaxSum)
			}
			if len(gotPath) != tt.wantPathLen {
				t.Errorf("solve() path length = %v, want %v", len(gotPath), tt.wantPathLen)
			}
		})
	}
}
