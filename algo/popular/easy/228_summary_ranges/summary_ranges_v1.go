package _28_summary_ranges

import "strconv"

func summaryRangesV1(nums []int) []string {
	/*
		METHOD: Loop
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	result := []string{}
	symbol := ""

	for i := 0; i < len(nums); {
		symbol = symbol + strconv.Itoa(nums[i])
		j := i + 1

		for ; j < len(nums) && nums[j]-nums[j-1] == 1; j++ {
		}

		if j-1 > i {
			symbol = symbol + "->" + strconv.Itoa(nums[j-1])
		}

		result = append(result, symbol)
		symbol = ""
		i = j
	}

	return result
}
