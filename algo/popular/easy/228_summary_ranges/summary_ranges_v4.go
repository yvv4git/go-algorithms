package _28_summary_ranges

import "fmt"

func summaryRangesV4(nums []int) []string {
	var result []string
	var cnt int

	for i, num := range nums {
		nextNum := num + 2
		if i != len(nums)-1 {
			nextNum = nums[i+1]
		}

		if nextNum-num <= 1 {
			cnt++

			continue
		}

		startNum := num - cnt
		if cnt == 0 {
			result = append(result, fmt.Sprintf("%d", num))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", startNum, num))
			cnt = 0
		}

	}

	return result
}
