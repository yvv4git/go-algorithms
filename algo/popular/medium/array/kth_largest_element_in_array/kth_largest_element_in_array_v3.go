package kth_largest_element_in_array

func findKthLargestV3(nums []int, k int) int {
	/*
		METHOD: Quick selection
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	n := len(nums)
	var quickSelect func(int, int, int) int

	quickSelect = func(left, right, k int) int {
		mid := (left + right) / 2
		pivot := nums[mid]
		nums[mid], nums[right-1] = nums[right-1], nums[mid]
		idx := left
		for i := left; i < right; i++ {
			if nums[i] < pivot {
				nums[i], nums[idx] = nums[idx], nums[i]
				idx += 1
			}
		}

		nums[idx], nums[right-1] = nums[right-1], nums[idx]
		if idx == right-k {
			return nums[idx]
		} else if idx < right-k {
			return quickSelect(idx+1, right, k)
		} else {
			return quickSelect(left, idx, k-(right-idx))
		}
	}

	return quickSelect(0, n, k)
}
