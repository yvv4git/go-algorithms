package power_of_four

func isPowerOfFourV2(n int) bool {
	/*
		Method: Bitwise
		Time complexity: O(1)
		Space complexity: O(1)
	*/
	return n != 0 && n&(n-1) == 0 && n&0xAAAAAAAA == 0
}
