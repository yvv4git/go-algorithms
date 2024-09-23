//go:build ignore

package main

import "sort"

func largestDivisibleSubsetV3(nums []int) []int {
	sort.Ints(nums)
	ans := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = append(ans[i], nums[i])
		for j := i + 1; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 && len(ans[j]) < len(ans[i]) {
				ans[j] = make([]int, len(ans[i]))
				copy(ans[j], ans[i])
			}
		}
	}

	var index int
	for i := range ans {
		if len(ans[i]) > len(ans[index]) {
			index = i
		}
	}
	return ans[index]
}
