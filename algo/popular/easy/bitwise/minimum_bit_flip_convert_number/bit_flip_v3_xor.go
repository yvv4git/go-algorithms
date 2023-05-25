package minimum_bit_flip_convert_number

func minBitFlipsV3(start int, goal int) int {
	ones := start ^ goal

	flipsCount := 0
	for ; ones != 0; ones &= ones - 1 {
		flipsCount++
	}

	return flipsCount
}
