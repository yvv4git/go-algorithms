package main

import "testing"

func Test_searchV2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1: Target is in the array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: true,
		},
		{
			name: "Test Case 2: Target is not in the array",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: false,
		},
		{
			name: "Test Case 3: Array contains duplicates and target is in the array",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 0,
			},
			want: true,
		},
		{
			name: "Test Case 4: Array contains duplicates and target is not in the array",
			args: args{
				nums:   []int{1, 0, 1, 1, 1},
				target: 2,
			},
			want: false,
		},
		{
			name: "Test Case 5: Array is not rotated",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 3,
			},
			want: true,
		},
		{
			name: "Test Case 6: Array contains only one element and it is the target",
			args: args{
				nums:   []int{5},
				target: 5,
			},
			want: true,
		},
		{
			name: "Test Case 7: Array contains only one element and it is not the target",
			args: args{
				nums:   []int{5},
				target: 3,
			},
			want: false,
		},
		{
			name: "Test Case 8: Array is empty",
			args: args{
				nums:   []int{},
				target: 3,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchV2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
