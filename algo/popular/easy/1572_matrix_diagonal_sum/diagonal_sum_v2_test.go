package main

import "testing"

func Test_diagonalSumV2(t *testing.T) {
	type args struct {
		mat [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				mat: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: 25,
		},
		{
			name: "CASE-2",
			args: args{
				mat: [][]int{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
				},
			},
			want: 8,
		},
		{
			name: "CASE-3",
			args: args{
				mat: [][]int{
					{5},
				},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalSumV2(tt.args.mat); got != tt.want {
				t.Errorf("diagonalSumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
