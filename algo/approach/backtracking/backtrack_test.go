package main

import "testing"

func Test_backtrack(t *testing.T) {
	type args struct {
		candidates []int
		target     int
		temp       []int
		result     *[][]int
	}

	tests := []struct {
		name     string
		args     args
		expected [][]int
	}{
		{
			name: "Test Case 1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
				temp:       []int{},
				result:     &[][]int{},
			},
			expected: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "Test Case 2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
				temp:       []int{},
				result:     &[][]int{},
			},
			expected: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "Test Case 3",
			args: args{
				candidates: []int{2},
				target:     1,
				temp:       []int{},
				result:     &[][]int{},
			},
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			backtrack(tt.args.candidates, tt.args.target, tt.args.temp, tt.args.result)
		})
	}
}
