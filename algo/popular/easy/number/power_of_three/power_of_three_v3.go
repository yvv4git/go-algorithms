package power_of_three

func isPowerOfThreeV3(n int) bool {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(1)
		Space complexity: O(1)
	*/
	if n == 1 {
		return true
	}

	if n < 1 || n%3 != 0 {
		return false
	}

	return isPowerOfThreeV3(n / 3)
}
