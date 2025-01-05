package main

func kidsWithCandiesMapV2(candies []int, extraCandies int) []bool {
	// First pass - find max value
	maxCandy := make(map[string]int)
	for _, v := range candies {
		if v > maxCandy["max"] {
			maxCandy["max"] = v
		}
	}

	// Second pass - compare with max
	result := make([]bool, len(candies))
	for i, v := range candies {
		if v+extraCandies >= maxCandy["max"] {
			result[i] = true
		}
	}

	return result
}
