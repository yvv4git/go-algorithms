package power_of_four

func isPowerOfFourV1(n int) bool {
	/*
		METHOD: Recursion
		Time complexity: ???
		Space complexity: O(1)
	*/
	if (n == 1) || (n > 0 && n&(n-1) == 0 && isPowerOfFourV1(n/4)) {
		return true
	}

	return false
}
