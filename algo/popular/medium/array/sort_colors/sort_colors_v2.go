package sort_colors

func sortColorsV2(nums []int) {
	/*
		METHOD: Counting sort.
		TIME COMPLEXITY: O(n+k), where n - length of nums, k - max number in nums (its at worst 2, so the complexity is basically can be considered O(n)).
		Space complexity: O(n+k) - same logic here.
	*/
	c := make([]int, max(nums)+1)
	for _, n := range nums {
		c[n]++
	}

	for i := 1; i < len(c); i++ {
		c[i] += c[i-1]
	}

	sorted := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		sorted[c[nums[i]]-1] = nums[i]
		c[nums[i]]--
	}

	copy(nums, sorted)
}

func max(nums []int) int {
	m := nums[0]
	for _, n := range nums {
		if n > m {
			m = n
		}
	}
	return m
}
