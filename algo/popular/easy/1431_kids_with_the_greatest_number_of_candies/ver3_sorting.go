package main

import "sort"

func kidsWithCandiesSortingV3(candies []int, extraCandies int) []bool {
	// Create a copy of candies array to sort
	sortedCandies := make([]int, len(candies))
	copy(sortedCandies, candies)

	// Sort in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(sortedCandies)))

	// Get maximum value (first element after sorting)
	maxCandy := sortedCandies[0]

	// Compare each original value + extraCandies with max
	result := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= maxCandy {
			result[i] = true
		}
	}

	return result
}
