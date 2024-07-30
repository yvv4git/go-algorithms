package main

import "testing"

func Test_searchMatrixV2(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Case 1: Element exists in the matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "Case 2: Element does not exist in the matrix",
			args: args{
				matrix: [][]int{
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 20,
			},
			want: false,
		},
		//{
		//	name: "Case 3: Empty matrix",
		//	args: args{
		//		matrix: [][]int{},
		//		target: 1,
		//	},
		//	want: false,
		//},
		{
			name: "Case 4: Single element matrix, element exists",
			args: args{
				matrix: [][]int{
					{1},
				},
				target: 1,
			},
			want: true,
		},
		{
			name: "Case 5: Single element matrix, element does not exist",
			args: args{
				matrix: [][]int{
					{1},
				},
				target: 2,
			},
			want: false,
		},
		{
			name: "Case 6: Large matrix, element exists",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				target: 13,
			},
			want: true,
		},
		{
			name: "Case 7: Large matrix, element does not exist",
			args: args{
				matrix: [][]int{
					{1, 2, 3, 4, 5},
					{6, 7, 8, 9, 10},
					{11, 12, 13, 14, 15},
					{16, 17, 18, 19, 20},
					{21, 22, 23, 24, 25},
				},
				target: 26,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrixV2(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrixV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
