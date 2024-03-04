package main

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "No obstacles",
			args: args{
				obstacleGrid: [][]int{
					{0, 0, 0},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: 6,
		},
		{
			name: "One obstacle",
			args: args{
				obstacleGrid: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "All obstacles",
			args: args{
				obstacleGrid: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: 0,
		},
		{
			name: "One row",
			args: args{
				obstacleGrid: [][]int{
					{0, 1, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "One column",
			args: args{
				obstacleGrid: [][]int{
					{0},
					{0},
					{1},
					{0},
				},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
