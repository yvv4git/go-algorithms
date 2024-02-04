package sort_colors

func sortColorsV6(nums []int) {
	/*
		METHOD: snowball algorithm.
		Time: O(n+n) = O(2n) = O(n)
		Space: O(1)
	*/
	// Ex: nums = []int{2, 0, 1, 0, 2, 1}

	// Time: O(n)
	moveToEnd(nums, 2)
	// Ex:nums = []int{0, 1, 0, 1, 2, 2}

	// Time: O(n)
	moveToBeginning(nums, 0)
	// Ex: nums = []int{0, 0, 1, 1, 2, 2}
}

func moveToEnd(nums []int, val int) {
	var snowballSize int = 0

	// Time: O(n)
	for idx, num := range nums {
		if num == val {
			snowballSize++
		} else if snowballSize > 0 {
			nums[idx], nums[idx-snowballSize] = val, num
		}
	}
}

func moveToBeginning(nums []int, val int) {
	var snowballSize int = 0

	// Time: O(n)
	for i := len(nums) - 1; i >= 0; i-- {
		if num := nums[i]; num == val {
			snowballSize++
		} else if snowballSize > 0 {
			nums[i], nums[i+snowballSize] = val, num
		}
	}
}
