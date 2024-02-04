package power_of_three

func isPowerOfThreeV1(n int) bool {
	/*
		TIME COMPLEXITY: O(1)
		SPACE COMPLEXITY: O(1)
	*/
	return n > 0 && 1162261467%n == 0
}
