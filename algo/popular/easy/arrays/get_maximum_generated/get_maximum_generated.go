package get_maximum_generated

func getMaximumGenerated(n int) int {
	result := 0
	nums := make([]int, n+1)

	if n >= 0 {
		nums[0] = 0
	}

	if n >= 1 {
		nums[1] = 1
		result = 1
	}

	for i := 2; i <= n; i++ {
		t := i / 2
		if i%2 == 0 {
			nums[i] = nums[t]
		} else {
			nums[i] = nums[t] + nums[t+1]
		}

		if nums[i] > result {
			result = nums[i]
		}
	}

	return result
}
