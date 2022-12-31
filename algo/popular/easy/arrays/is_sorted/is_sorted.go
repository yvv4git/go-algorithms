package is_sorted

// IsSorted - check array is sorted and oriented.
func IsSorted(nums []int) bool {
	shift := 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			if shift == 0 {
				shift = i
			} else { // array was not sorted because
				return false // we found 2nd place with decreasing value
			}
		}
	}

	if shift > 0 { // check that the last and the first values are in order
		return nums[len(nums)-1] <= nums[0]
	}

	return true // zero shift means no rotation was applied
}
