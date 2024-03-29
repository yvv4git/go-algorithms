package kth_largest_element_in_array

func findKthLargestV1(nums []int, k int) int {
	/*
		METHOD: Hash
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(max(nums) - min(nums)+1)
	*/
	const MaxInt = 2147483647
	const MinInt = -2147483648

	// TIME COMPLEXITY: O(n) = O(nums)
	maxValue, minValue := MinInt, MaxInt
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}

		if nums[i] < minValue {
			minValue = nums[i]
		}
	}

	// SPACE COMPLEXITY: O(max(nums) - min(nums)+1)
	count := make([]int, maxValue-minValue+1)
	for i := 0; i < len(nums); i++ {
		count[nums[i]-minValue]++
	}

	// TIME COMPLEXITY: O(n) = O(count)
	var result int
	for i := len(count) - 1; i >= 0 && k > 0; i-- {
		if count[i] > 0 {
			result = i + minValue
			k -= count[i]
		}
	}

	return result
}
