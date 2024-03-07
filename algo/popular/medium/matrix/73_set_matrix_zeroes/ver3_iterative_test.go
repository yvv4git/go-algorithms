package main

import "testing"

func Test_setZeroesV3(t *testing.T) {
	type args struct {
		matrix [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 0, 1},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				matrix: [][]int{
					{0, 1, 2, 0},
					{3, 4, 5, 2},
					{1, 3, 1, 5},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				matrix: [][]int{
					{1, 0, 3},
				},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "Test Case 4",
			args: args{
				matrix: [][]int{
					{1},
					{0},
					{3},
				},
			},
			want: [][]int{
				{0},
				{0},
				{0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setZeroesV3(tt.args.matrix)
		})
	}
}
