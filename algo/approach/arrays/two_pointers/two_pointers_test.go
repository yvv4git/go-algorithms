package main

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				arr: []int{1, 4, 2, 10, 23, 3, 1, 0, 20},
				k:   4,
			},
			want: 39,
		},
		{
			name: "Test Case 2",
			args: args{
				arr: []int{100, 200, 300, 400},
				k:   2,
			},
			want: 700,
		},
		{
			name: "Test Case 3",
			args: args{
				arr: []int{2, 3},
				k:   3,
			},
			want: -1, // Since k is greater than the length of the array
		},
		{
			name: "Test Case 4",
			args: args{
				arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				k:   5,
			},
			want: 40, // The maximum sum of a subarray of size 5 is 6 + 7 + 8 + 9 + 10 = 40
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
