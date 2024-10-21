package main

import "sort"

func distributeCandiesV4(candyType []int) int {
	/*
		METHOD: Sort
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: O(1)
	*/
	sort.Ints(candyType)
	numUniqueCandies := 1

	for i := 1; i < len(candyType); i++ {
		if candyType[i] != candyType[i-1] {
			numUniqueCandies++
		}
	}

	maxCandies := len(candyType) / 2
	if numUniqueCandies > maxCandies {
		return maxCandies
	}

	return numUniqueCandies
}
