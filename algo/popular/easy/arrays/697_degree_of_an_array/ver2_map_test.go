package main

import "testing"

func Test_findShortestSubArrayV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{nums: []int{1, 2, 2, 3, 1}},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{nums: []int{1, 2, 2, 3, 1, 4, 2}},
			want: 6,
		},
		{
			name: "Test Case 3",
			args: args{nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3}},
			want: 3,
		},
		{
			name: "Test Case 4",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "Test Case 5",
			args: args{nums: []int{1, 1, 1, 1, 1}},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArrayV2(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArrayV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
