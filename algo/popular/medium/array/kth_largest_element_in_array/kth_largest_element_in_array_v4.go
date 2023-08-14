package kth_largest_element_in_array

func findKthLargestV4(nums []int, k int) int {
	/*
		Method: Hash, memory
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	left := []int{}  // Hash left
	right := []int{} // Hash right
	middle := len(nums) / 2
	pivot := nums[middle]

	for i, item := range nums {
		if i == middle {
			continue
		}
		if item >= pivot {
			right = append(right, item)
		} else {
			left = append(left, item)
		}
	}

	if len(right) == k-1 {
		return pivot
	} else if len(right) > k-1 {
		return findKthLargestV4(right, k)
	} else {
		return findKthLargestV4(left, k-len(right)-1)
	}
}
