package sum_all_subset_xor_totals

import (
	"fmt"
	"testing"
)

func Test_subsetXORSumV3(t *testing.T) {
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
			if got := subsetXORSumV3(tt.args.nums); got != tt.want {
				t.Errorf("subsetXORSumV3() = %v, want %v", got, tt.want)
			}
		})
	}
}

// -------------------------
func xorSumResearch(list []int) int {
	/*
		Получается, в эту функцию передается список всех возможных комбинаций.
		И она возвращает результат. Если посмотреть лог, то видно, что туда попали все возможные сочетания(не комбинации).
		list: []int{} -> 9
		list: []int{5} -> 5
		list: []int{5, 1} -> 4
		list: []int{5, 1, 6} -> 2
		list: []int{5, 6} -> 3
		list: []int{1} -> 1
		list: []int{1, 6} -> 7
		list: []int{6} -> 6
	*/
	fmt.Printf("list: %#v ", list)
	if len(list) == 0 {
		fmt.Printf("-> %#v \n", 0)
		return 0
	}
	res := list[0]
	for i := 1; i < len(list); i++ {
		res ^= list[i]
	}
	fmt.Printf("-> %#v \n", res)
	return res
}

func backTrackResearch(nums []int, pos int, list []int, res *int) {
	fmt.Printf("backTrack(nums=%v, pos=%d, list=%v, res=%v) \n", nums, pos, list, *res)
	*res += xorSumResearch(list)
	fmt.Printf(" -> %v \n", *res)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backTrackResearch(nums, i+1, list, res)
		fmt.Printf("--> list-1=%v \n", list)
		list = list[:len(list)-1]
		fmt.Printf("--> list-2=%v \n", list)
	}
}

func TestResearchM3Ex01(t *testing.T) {
	res := 0
	list := make([]int, 0)
	backTrackResearch([]int{5, 1, 6}, 0, list, &res)

	t.Logf("Result: %v", res)
}
