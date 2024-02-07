package main

import "testing"

func Test_nextPermutationV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test Case 1: Simple case",
			args: args{nums: []int{1, 2, 3}},
			want: []int{1, 3, 2},
		},
		{
			name: "Test Case 2: Reversed case",
			args: args{nums: []int{3, 2, 1}},
			want: []int{1, 2, 3},
		},
		{
			name: "Test Case 3: Single element",
			args: args{nums: []int{1}},
			want: []int{1},
		},
		{
			name: "Test Case 4: Empty array",
			args: args{nums: []int{}},
			want: []int{},
		},
		{
			name: "Test Case 5: Case with same elements",
			args: args{nums: []int{1, 1, 1}},
			want: []int{1, 1, 1},
		},
		{
			name: "Test Case 6: Case with multiple same elements",
			args: args{nums: []int{1, 1, 2, 2}},
			want: []int{1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutationV2(tt.args.nums)
		})
	}
}
