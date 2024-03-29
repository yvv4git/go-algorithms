package squares_sorted_array

import "sort"

func sortedSquaresV1(nums []int) []int {
	/*
		METHOD: Use default go sort.
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: ???
	*/
	for idx, num := range nums {
		num *= num
		nums[idx] = num
	}
	sort.Ints(nums) // This is the reason why it's O(nlogn)

	return nums
}
