package count_number_of_consistent_strings

func countConsistentStringsV2(allowed string, words []string) int {
	/*
		METHOD: Bitwise
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	result := 0
	bitMask := 0
	bitMask = getAllowedBitMask(allowed)
	for _, word := range words {
		check := getAllowedBitMask(word)
		if (check | bitMask) == bitMask {
			result++
		}
	}

	return result
}

func getAllowedBitMask(word string) int {
	var bitMask int = 0

	for _, char := range word {
		bitMask = bitMask | (1 << (char - 'a'))
	}

	return bitMask
}
