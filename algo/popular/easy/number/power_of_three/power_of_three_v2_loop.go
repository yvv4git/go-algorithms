package power_of_three

func isPowerOfThreeV2(n int) bool {
	/*
		METHOD: For loop
		TIME COMPLEXITY: O(1)
		Space complexity: O(1)
	*/
	if n < 1 {
		return false
	}

	for n != 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}

	return true
}
