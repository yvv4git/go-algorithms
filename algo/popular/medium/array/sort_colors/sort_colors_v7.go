package sort_colors

func sortColorsV7(nums []int) {
	/*
		Method: One-Pass Three-Pointers
		Time: O(n)
		Space: O(1)
	*/
	// Time: O(n)
	for low, high, ptr := 0, len(nums)-1, 0; ptr <= high; {
		num := nums[ptr]
		if num == 0 {
			nums[ptr], nums[low] = nums[low], nums[ptr]
			ptr++
			low++
		} else if num == 1 {
			ptr++
		} else /* num == 2 */ {
			nums[ptr], nums[high] = nums[high], nums[ptr]
			high--
		}
	}
}
