package _28_summary_ranges

import "strconv"

func summaryRangesV3(nums []int) []string {
	/*
		Method: Recursive
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	if len(nums) == 0 {
		return []string{}
	}

	var sol []string
	ptr := nums[0]
	streak := nums[0]
	i := 1
	for i < len(nums) && nums[i] == streak+1 { //build your streak
		streak++
		i++
	}

	// at this point, i points to "out of range" or the number that is after the streak
	// in the following step, append the summarized range we calculated earlier. Don't worry about recursion yet
	if ptr != streak {
		start := strconv.Itoa(ptr)
		end := strconv.Itoa(streak)
		sol = append(sol, start+"->"+end)
	} else {
		start := strconv.Itoa(ptr)
		sol = append(sol, start)
	}

	sol = append(sol, summaryRangesV3(nums[i:])...) //recursively send next number in nums for summarizing
	return sol
}
