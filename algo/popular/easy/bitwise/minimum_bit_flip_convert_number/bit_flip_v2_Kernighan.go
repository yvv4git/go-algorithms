package minimum_bit_flip_convert_number

func minBitFlipsV2(start int, goal int) int {
	/*
		Method: Kernighan algorithm
		Time complexity: O(log(n))
		Space complexity: O(1)
	*/
	result := 0
	for x := start ^ goal; x != 0; x &= x - 1 {
		result++
	}

	return result
}
