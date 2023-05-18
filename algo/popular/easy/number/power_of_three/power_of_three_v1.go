package power_of_three

func isPowerOfThreeV1(n int) bool {
	/*
		Time complexity: O(1)
		Space complexity: O(1)
	*/
	return n > 0 && 1162261467%n == 0
}
