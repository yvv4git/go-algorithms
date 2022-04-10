package numbersdisappearedinarray

func findDisappearedNumbers(nums []int) []int {
	abs := func(n int) int {
		if n > 0 {
			return n
		}

		return -n
	}

	ans := []int{}

	// Хитрая закономерность.
	for _, v := range nums {
		idx := abs(v) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		} else {
		}
	}

	for i, v := range nums {
		if v > 0 {
			ans = append(ans, i+1)
		}
	}

	return ans
}
