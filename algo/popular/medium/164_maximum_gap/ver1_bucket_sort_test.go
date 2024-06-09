package main

import "testing"

func Test_maximumGap(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Regular case with a maximum gap",
			args: args{nums: []int{3, 6, 9, 1}},
			want: 3, // The maximum gap is between 9 and 6 which is 3
		},
		{
			name: "Test Case 2: Array with two elements",
			args: args{nums: []int{10, 20}},
			want: 10, // The maximum gap is between 20 and 10 which is 10
		},
		{
			name: "Test Case 3: Array with equal elements",
			args: args{nums: []int{5, 5, 5, 5}},
			want: 0, // All elements are the same, so the maximum gap is 0
		},
		{
			name: "Test Case 4: Array with one element",
			args: args{nums: []int{100}},
			want: 0, // Only one element, so the maximum gap is 0
		},
		{
			name: "Test Case 5: Empty array",
			args: args{nums: []int{}},
			want: 0, // Empty array, so the maximum gap is 0
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGap(tt.args.nums); got != tt.want {
				t.Errorf("maximumGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
