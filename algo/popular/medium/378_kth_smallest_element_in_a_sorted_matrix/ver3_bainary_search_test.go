package main

import "testing"

func Test_kthSmallestV3(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}},
				k:      8,
			},
			want: 13,
		},
		{
			name: "Example 2",
			args: args{
				matrix: [][]int{{-5}},
				k:      1,
			},
			want: -5,
		},
		{
			name: "Example 3",
			args: args{
				matrix: [][]int{{1, 2}, {1, 3}},
				k:      2,
			},
			want: 1,
		},
		{
			name: "Example 4",
			args: args{
				matrix: [][]int{{1, 3, 5}, {6, 7, 12}, {11, 14, 14}},
				k:      6,
			},
			want: 11,
		},
		{
			name: "Example 5",
			args: args{
				matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				k:      5,
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestV3(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
