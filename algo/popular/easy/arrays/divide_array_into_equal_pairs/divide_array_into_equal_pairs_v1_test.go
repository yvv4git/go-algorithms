package divide_array_into_equal_pairs

import "testing"

func Test_divideArrayV1(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want bool
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{3, 2, 3, 2, 2, 2},
			},
			want: true,
			desc: `Explanation: There are 6 elements in nums, so they should be divided into 6 / 2 = 3 pairs.
					If nums is divided into the pairs (2, 2), (3, 3), and (2, 2), it will satisfy all the conditions.`,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
			desc: `Explanation: There is no way to divide nums into 4 / 2 = 2 pairs such that the pairs satisfy every condition.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divideArrayV1(tt.args.nums); got != tt.want {
				t.Errorf("divideArrayV1() = %v, want %v", got, tt.want)
			}
		})
	}
}
