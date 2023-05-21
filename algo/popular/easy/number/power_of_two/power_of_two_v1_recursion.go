package power_of_two

func isPowerOfTwoV1(n int) bool {
	/*
		Method: Recursion
		Time complexity: O(1)
		Space complexity: O(1)
	*/
	return power2(float64(n))
}

func power2(n float64) bool {
	if n == 1 {
		return true
	}
	if n < 2 {
		return false
	}

	return power2(n / 2)
}
