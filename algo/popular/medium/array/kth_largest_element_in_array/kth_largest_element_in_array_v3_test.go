package kth_largest_element_in_array

import "testing"

func Test_findKthLargestV3(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{3, 2, 1, 5, 6, 4},
				k:    2,
			},
			want: 5,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
				k:    4,
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargestV3(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargestV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
