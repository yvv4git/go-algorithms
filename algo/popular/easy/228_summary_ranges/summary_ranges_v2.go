package _28_summary_ranges

import "fmt"

func summaryRangesV2(nums []int) []string {
	/*
		METHOD: Window
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	if len(nums) == 0 {
		return []string{}
	}

	var ans []string
	left := nums[0] // Start with the first value
	prev := nums[0] // Start with the first value
	for i := 1; i < len(nums); i++ {
		if nums[i]-prev != 1 {
			// We've reached right border of current range
			ans = append(ans, strProcessing(left, prev))
			// Save the begining of new range
			left, prev = nums[i], nums[i]
		} else {
			// Expand our window further and update prev value
			prev = nums[i]
		}
	}

	// Don't forget to append last found range
	return append(ans, strProcessing(left, prev))
}

func strProcessing(l, r int) string {
	if l == r {
		return fmt.Sprint(l)
	}

	return fmt.Sprint(l, "->", r)
}
