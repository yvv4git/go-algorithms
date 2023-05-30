package sum_all_subset_xor_totals

import (
	"fmt"
	"testing"
)

func Test_subsetXORSumV2(t *testing.T) {
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
			if got := subsetXORSumV2(tt.args.nums); got != tt.want {
				t.Errorf("subsetXORSumV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
// ---> Example-1
func TestResearchMethodV1Ex01(t *testing.T) {
	n := []int{5, 1, 6}
	total := 0
	for _, num := range n {
		total = total ^ num
	}

	t.Logf("Total: %d", total) // 2
}

// ---> Example-2
func TestResearchMethodV1Ex02(t *testing.T) {
	n := []int{2, 5, 6}
	total := 0
	for _, num := range n {
		total = total ^ num
	}

	t.Logf("Total: %d", total) // 1
}

// ---> Example-3
func TestResearchMethodV1Ex03(t *testing.T) {
	total := 0
	total = total ^ 5

	t.Logf("Total: %d", total) // 5
}

// ---> Example-4
func dfsResearch(idx int, nums []int, curXor int, level int, name string) int {
	/*
		Explanation: The 8 subsets of [5,1,6] are:
		- The empty subset has an XOR total of 0.
		- [5] has an XOR total of 5.
		- [1] has an XOR total of 1.
		- [6] has an XOR total of 6.
		- [5,1] has an XOR total of 5 XOR 1 = 4.
		- [5,6] has an XOR total of 5 XOR 6 = 3.
		- [1,6] has an XOR total of 1 XOR 6 = 7.
		- [5,1,6] has an XOR total of 5 XOR 1 XOR 6 = 2.
		0 + 5 + 1 + 6 + 4 + 3 + 7 + 2 = 28
	*/
	//fmt.Printf("[%d] dfs(%d, %v, %d) \n", level, idx, nums, curXor)
	if len(nums) == 0 || idx == len(nums) {
		fmt.Printf("[%d][%s] idx=%d curXor=%d nums=%v \n", level, name, idx, curXor, nums)
		return curXor
	}

	withNum := dfsResearch(idx+1, nums, curXor^nums[idx], level+1, "withNum")
	withoutNum := dfsResearch(idx+1, nums, curXor, level+1, "withoutNum")

	//fmt.Printf("[%d] withNum(%d) + withoutNum(%d) = %d  idx=%d curXor=%d nums=%v \n", level, withNum, withoutNum, withNum+withoutNum, idx, curXor, nums)
	return withNum + withoutNum
}

func TestResearchMethodV1Ex04(t *testing.T) {
	result := dfsResearch(0, []int{5, 1, 6}, 0, 0, "root")
	t.Logf("Result: %d", result)
}
