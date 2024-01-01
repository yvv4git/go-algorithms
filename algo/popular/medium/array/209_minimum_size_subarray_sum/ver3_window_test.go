package _09_minimum_size_subarray_sum

import "testing"

func Test_minSubArrayLenV3(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				target: 7,
				nums:   []int{2, 3, 1, 2, 4, 3},
			},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{
				target: 4,
				nums:   []int{1, 4, 4},
			},
			want: 1,
		},
		{
			name: "Test Case 3",
			args: args{
				target: 11,
				nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			},
			want: 0,
		},
		{
			name: "Test Case 4",
			args: args{
				target: 15,
				nums:   []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenV3(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLenV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
