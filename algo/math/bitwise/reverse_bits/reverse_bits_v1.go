package reverse_bits

func reverseBitsV1(n int) int {
	/*
		TIME COMPLEXITY: O(num), where num is the number of bits in binary representation of n
		SPACE COMPLEXITY: O(1)
	*/
	result := 0

	for n > 0 {
		result <<= 1 // Left shift by 1
		if n&1 == 1 {
			result = result ^ 1
		}

		n >>= 1 // Right shift by 1
	}

	return result
}
