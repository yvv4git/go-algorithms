package max_consecutive_one

func findMaxConsecutiveOnes(nums []int) int {
	c, tmp := 0, 0
	for _, v := range nums {
		if v == 1 {
			tmp++
			c = max(c, tmp)
		} else {
			tmp = 0
		}
	}
	return max(c, tmp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
