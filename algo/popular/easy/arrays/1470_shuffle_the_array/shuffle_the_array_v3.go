package _470_shuffle_the_array

func shuffleV3(nums []int, n int) []int {
	/*
		METHOD: Iterate
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)
	*/
	var result []int

	for i := 0; i != n; i++ {
		result = append(result, nums[i], nums[n+i])
	}

	return result
}
