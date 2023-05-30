package sum_all_subset_xor_totals

import (
	"fmt"
	"testing"
)

func Test_subsetXORSumV1(t *testing.T) {
	type args struct {
		nums []int
	}

	tests := []struct {
		name string
		args args
		want int
		desc string
	}{
		{
			name: "CASE-1",
			args: args{
				nums: []int{1, 3},
			},
			want: 6,
			desc: `Explanation: The 4 subsets of [1,3] are:
					- The empty subset has an XOR total of 0.
					- [1] has an XOR total of 1.
					- [3] has an XOR total of 3.
					- [1,3] has an XOR total of 1 XOR 3 = 2.
					0 + 1 + 3 + 2 = 6`,
		},
		{
			name: "CASE-2",
			args: args{
				nums: []int{5, 1, 6},
			},
			want: 28,
			desc: `Explanation: The 8 subsets of [5,1,6] are:
					- The empty subset has an XOR total of 0.
					- [5] has an XOR total of 5.
					- [1] has an XOR total of 1.
					- [6] has an XOR total of 6.
					- [5,1] has an XOR total of 5 XOR 1 = 4.
					- [5,6] has an XOR total of 5 XOR 6 = 3.
					- [1,6] has an XOR total of 1 XOR 6 = 7.
					- [5,1,6] has an XOR total of 5 XOR 1 XOR 6 = 2.
					0 + 5 + 1 + 6 + 4 + 3 + 7 + 2 = 28`,
		},
		{
			name: "CASE-3",
			args: args{
				nums: []int{3, 4, 5, 6, 7, 8},
			},
			want: 480,
			desc: `Explanation: The sum of all XOR totals for every subset is 480.`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsetXORSumV1(tt.args.nums); got != tt.want {
				t.Errorf("subsetXORSumV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func dfsResearchV1(nums []int, cur int, res *int) {
	fmt.Printf("nums=%v  cur=%d\n", nums, cur)
	if len(nums) == 0 {
		*res += cur
		return
	}

	dfsResearchV1(nums[1:], cur^nums[0], res)
	dfsResearchV1(nums[1:], cur, res)
}

func TestResearchM1V1(t *testing.T) {
	res := 0
	dfsResearchV1([]int{1, 2, 3}, 0, &res)
}
