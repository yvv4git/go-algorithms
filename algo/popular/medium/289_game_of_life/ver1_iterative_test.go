package main

import "testing"

func Test_gameOfLife(t *testing.T) {
	type args struct {
		board [][]int
	}

	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Empty board",
			args: args{
				board: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "Single cell dies",
			args: args{
				board: [][]int{
					{1},
				},
			},
			want: [][]int{
				{0},
			},
		},
		{
			name: "Block pattern remains stable",
			args: args{
				board: [][]int{
					{1, 1},
					{1, 1},
				},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			name: "Blinker pattern oscillates",
			args: args{
				board: [][]int{
					{0, 1, 0},
					{0, 1, 0},
					{0, 1, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 1, 1},
				{0, 0, 0},
			},
		},
		{
			name: "Glider pattern moves",
			args: args{
				board: [][]int{
					{0, 1, 0, 0},
					{0, 0, 1, 0},
					{1, 1, 1, 0},
					{0, 0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{0, 1, 1, 0},
				{0, 1, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.args.board)
		})
	}
}
