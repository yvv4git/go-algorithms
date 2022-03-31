package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 1
	for i < len(nums) {
		if nums[i] == val {
			for j < len(nums) && nums[j] == val {
				j++
			}
			if j >= len(nums) {
				return i
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
		j++
	}
	return i
}
