package _464_max_product_of_two_elem_of_arr

import "sort"

func maxProductV1(nums []int) int {
	/*
		METHOD: Use default sort.
		TIME COMPLEXITY: O(n log(n))
		Space complexity: O(1)
	*/
	sort.Ints(nums)
	i := len(nums) - 1
	j := len(nums) - 2

	return (nums[i] - 1) * (nums[j] - 1)
}
