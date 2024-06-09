package _875_minimum_size_subarray_in_infinite_array

import "testing"

func Test_minSizeSubarrayV2(t *testing.T) {
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
			name: "Test Case 1",
			args: args{
				nums:   []int{1, 2, 3},
				target: 5,
			},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{
				nums:   []int{1, 1, 1, 2, 3},
				target: 4,
			},
			want: 2,
		},
		{
			name: "Test Case 3",
			args: args{
				nums:   []int{2, 4, 6, 8},
				target: 3,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSizeSubarrayV2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("minSizeSubarrayV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
