package _470_shuffle_the_array

func shuffleV1(nums []int, n int) []int {
	/*
		Method: Iterate in-place
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	i, j := n-1, len(nums)-1
	for i >= 0 {
		nums[j] <<= 10
		nums[j] |= nums[i]
		j--
		i--
	}

	i, j = 0, n
	for j < len(nums) {
		nums[i], nums[i+1] = nums[j]&1023, nums[j]>>10
		i += 2
		j++
	}

	return nums
}
