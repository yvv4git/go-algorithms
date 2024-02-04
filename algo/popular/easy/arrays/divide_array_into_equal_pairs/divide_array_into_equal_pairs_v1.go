package divide_array_into_equal_pairs

import "sort"

func divideArrayV1(nums []int) bool {
	/*
		METHOD: Use lib sort.
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)
	*/
	sort.Ints(nums) // TIME COMPLEXITY: O(n log(n))
	res := true

	// TIME COMPLEXITY: O(n)
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return false
		}
	}

	return res
}
