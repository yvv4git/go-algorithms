package main

func distributeCandiesV2(candyType []int) int {
	/*
		METHOD: Slice
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	uniqueCandies := []int{}
	for _, candy := range candyType {
		if !contains(uniqueCandies, candy) {
			uniqueCandies = append(uniqueCandies, candy)
		}
	}
	numUniqueCandies := len(uniqueCandies)
	if numUniqueCandies >= len(candyType)/2 {
		return len(candyType) / 2
	} else {
		return numUniqueCandies
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
