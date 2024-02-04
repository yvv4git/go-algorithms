package intersection_multiple_arrays

func intersectionV1(nums [][]int) []int {
	/*
		METHOD: Hash
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)
	*/
	const max = 1001 // From constraints

	hash := make([]int, max)
	for i := range nums {
		for j := range nums[i] {
			hash[nums[i][j]]++
		}
	}

	result := []int{}
	for i := range hash {
		if hash[i] == len(nums) {
			result = append(result, i)
		}
	}

	return result
}
