package sort_colors

func sortColorsV3(nums []int) {
	/*
		Method: Selection sort.
		Time complexity : O(n^2)
		Space complexity : O(1)
	*/
	for i := 0; i < len(nums)-1; i++ {
		pos := i
		for i := i + 1; i < len(nums); i++ {
			if nums[pos] > nums[i] {
				pos = i
			}
		}

		nums[pos], nums[i] = nums[i], nums[pos]
	}
}
