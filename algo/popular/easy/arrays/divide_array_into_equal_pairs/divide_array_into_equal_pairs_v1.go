package divide_array_into_equal_pairs

import "sort"

func divideArrayV1(nums []int) bool {
	/*
		Method: Use lib sort.
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	sort.Ints(nums) // Time complexity: O(n log(n))
	res := true

	// Time complexity: O(n)
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return false
		}
	}

	return res
}
