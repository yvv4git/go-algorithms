package _875_minimum_size_subarray_in_infinite_array

import "testing"

func Test_minSubArrayLenV1(t *testing.T) {
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
				nums:   []int{2, 3, 1, 2, 4, 3},
				target: 7,
			},
			want: 3,
		},
		{
			name: "Test Case 2",
			args: args{
				nums:   []int{1, 4, 4},
				target: 4,
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
				target: 11,
			},
			want: -1,
		},
		{
			name: "Test Case 4",
			args: args{
				nums:   []int{1, 2, 3},
				target: 5,
			},
			want: 2,
		},
		{
			name: "Test Case 5",
			args: args{
				nums:   []int{1, 1, 1, 2, 3},
				target: 4,
			},
			want: 3,
		},
		{
			name: "Test Case 6",
			args: args{
				nums:   []int{2, 4, 6, 8},
				target: 3,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenV1(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("minSubArrayLenV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
