package _427_number_of_common_factors

func commonFactorsV1(a int, b int) int {
	/*
		METHOD: Math
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	min := a
	if min > b {
		min = b
	}

	// Start from 2 because 1 is a default. Loop and check them until reaching the minimum number.
	result := 1
	for n := 2; n <= min; n++ {
		if a%n == 0 && b%n == 0 {
			result++
		}
	}

	return result
}
