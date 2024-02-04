package kth_largest_element_in_array

import "math"

func findKthLargestV5(nums []int, k int) int {
	/*
		METHOD: Counting sort with memory optimization
		TIME COMPLEXITY: O(n + m), where m = max - min
		Space complexity: O(m), where m = max - min
	*/
	min := math.MaxInt32
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	build := make([]int, max-min+1)
	for _, num := range nums {
		build[num-min]++
	}

	i := 0
	for j, bNum := range build {
		for k := 0; k < bNum; k++ {
			nums[i] = j + min
			i++
		}
	}

	return nums[len(nums)-k]
}
