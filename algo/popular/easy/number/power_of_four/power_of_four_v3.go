package power_of_four

const SecondBit int = 0x55555555 //  in binary ...1010101

func isPowerOfFourV3(n int) bool {
	/*
		METHOD: Bitwise
		Time complexity: O(1)
		Space complexity: O(1)

		Explanation:
		We can observe that every power of 4 has only 1 bit, and it's always in the odd(нечетное) positions. (1st, 3rd, 5th).
		4 ^ 0 = 1 -> 00000001
		4 ^ 1 = 4 -> 00000100
		4 ^ 2 = 16 -> 00010000
	*/
	return n > 0 && // Can't have 0 or negatives
		n&(n-1) == 0 && // Checking if only 1 bit is active
		n&SecondBit != 0 // Checking if the bit is in an odd position
}
