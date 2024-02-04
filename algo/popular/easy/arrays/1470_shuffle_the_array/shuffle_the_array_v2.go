package _470_shuffle_the_array

func shuffleV2(nums []int, n int) []int {
	/*
		METHOD: Iterate
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	result := []int{}

	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}

	return result
}
