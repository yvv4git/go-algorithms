package main

import "math/bits"

func distributeCandiesV3(candyType []int) int {
	/*
		METHOD: Bitset
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	bitset := 0
	for _, candy := range candyType {
		bitset |= 1 << candy
	}
	numUniqueCandies := bits.OnesCount(uint(bitset))
	maxCandies := len(candyType) / 2
	if numUniqueCandies > maxCandies {
		return maxCandies
	}

	return numUniqueCandies
}
