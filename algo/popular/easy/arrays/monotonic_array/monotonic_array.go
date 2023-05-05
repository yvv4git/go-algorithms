package monotonic_array

func isMonotonic(nums []int) bool {
	/*
		Time complexity : O(n)
		Space complexity : O(n)
	*/
	isIncreasing := true
	isDecreasing := true

	for i := 0; i < len(nums)-1; i++ {
		isIncreasing = isIncreasing && nums[i] <= nums[i+1]
		isDecreasing = isDecreasing && nums[i] >= nums[i+1]
	}

	return isIncreasing || isDecreasing
}
