package main

import "testing"

func Test_removeDuplicatesV2(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1: No duplicates",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "Test Case 2: All duplicates",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: 2,
		},
		{
			name: "Test Case 3: Mixed with duplicates",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 3},
			},
			want: 5,
		},
		{
			name: "Test Case 4: Empty array",
			args: args{
				nums: []int{},
			},
			want: 0,
		},
		{
			name: "Test Case 5: Single element array",
			args: args{
				nums: []int{1},
			},
			want: 1,
		},
		{
			name: "Test Case 6: Array with one duplicate",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: 3,
		},
		{
			name: "Test Case 7: Array with multiple duplicates",
			args: args{
				nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesV2(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicatesV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
