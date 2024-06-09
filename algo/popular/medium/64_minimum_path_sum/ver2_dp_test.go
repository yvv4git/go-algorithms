package main

import "testing"

func Test_minPathSumV2(t *testing.T) {
	type args struct {
		grid [][]int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				grid: [][]int{
					{1, 3, 1},
					{1, 5, 1},
					{4, 2, 1},
				},
			},
			want: 7, // Минимальная сумма пути должна быть 1->3->1->1->1 = 7
		},
		{
			name: "Test Case 2",
			args: args{
				grid: [][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			},
			want: 12, // Минимальная сумма пути должна быть 1->2->3->6 = 12
		},
		{
			name: "Test Case 3",
			args: args{
				grid: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			want: 5, // Минимальная сумма пути должна быть 1->1->1->1->1 = 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPathSumV2(tt.args.grid); got != tt.want {
				t.Errorf("minPathSumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
