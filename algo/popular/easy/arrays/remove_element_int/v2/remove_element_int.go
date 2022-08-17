package v2

func removeElement(nums []int, val int) int {
	count := 0

	for i, num := range nums {
		nums[i-count] = nums[i]
		if num == val {
			count += 1
		}
	}

	return len(nums) - count
}
