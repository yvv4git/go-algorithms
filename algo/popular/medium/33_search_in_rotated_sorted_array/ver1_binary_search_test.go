package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: Element exists in the array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "Test Case 2: Element does not exist in the array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Test Case 3: Array is empty",
			args: args{
				nums:   []int{},
				target: 5,
			},
			want: -1,
		},
		{
			name: "Test Case 4: Array has one element and it is the target",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: 0,
		},
		{
			name: "Test Case 5: Array has one element and it is not the target",
			args: args{
				nums:   []int{5},
				target: 3,
			},
			want: -1,
		},
		{
			name: "Test Case 6: Target is the first element",
			args: args{
				nums:   []int{3, 1},
				target: 3,
			},
			want: 0,
		},
		{
			name: "Test Case 7: Target is the last element",
			args: args{
				nums:   []int{5, 1, 3},
				target: 3,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
