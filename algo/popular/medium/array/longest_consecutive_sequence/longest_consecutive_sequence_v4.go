package longest_consecutive_sequence

func longestConsecutiveV4(nums []int) int {
	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}

	var globalLength int

	for _, num := range nums {
		// We'd like to start with the lowest value.
		// For example, if we have: 1, 2, 3, 4 then there is no point to start with 2, 3 or 4,
		// since the length will be lower than if we'd start with 1
		// therefore we skip these kind of values
		if set[num-1] {
			continue
		}

		localLength := 1
		for set[num+localLength] {
			localLength++
		}

		if localLength > globalLength {
			globalLength = localLength
		}
	}

	return globalLength
}
