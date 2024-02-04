package squares_sorted_array

func sortedSquaresV2(nums []int) []int {
	/*
		METHOD: Two pointers.
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: ???
	*/
	// STEP 1
	// First, we square all the nums
	for idx, num := range nums {
		num *= num
		nums[idx] = num
	}

	if len(nums) <= 1 {
		return nums
	}

	// STEP 2
	// After this, we move our two pointers where the smallest number is found
	// (near center, could be beginning or end)
	var ptr1, ptr2 int = 0, 1
	for ptr2 < len(nums)-1 && nums[ptr1] >= nums[ptr1+1] && nums[ptr2] >= nums[ptr2+1] {
		ptr1++
		ptr2++
	}

	// STEP 3
	// We add the smallest number to the result slice (array)
	//
	// If we add the value of the first pointer `ptr1`, we move it to left
	// Else if we add the value of the second pointer `ptr2`, we move it to the right
	//
	// If we are done with one pointer, to add values of the other pointers.
	// And we do this until we can't.
	result := make([]int, 0)
	for {
		if ptr1 < 0 && ptr2 >= len(nums) {
			return result
		} else if ptr1 < 0 {
			result = append(result, nums[ptr2])
			ptr2++
		} else if ptr2 >= len(nums) {
			result = append(result, nums[ptr1])
			ptr1--
		} else if nums[ptr1] <= nums[ptr2] {
			result = append(result, nums[ptr1])
			ptr1--
		} else {
			result = append(result, nums[ptr2])
			ptr2++
		}
	}

	return result
}
