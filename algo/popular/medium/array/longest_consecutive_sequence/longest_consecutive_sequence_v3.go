package longest_consecutive_sequence

func longestConsecutiveV3(nums []int) int {
	/*
		TIME COMPLEXITY: O(n).
		SPACE COMPLEXITY: O(n).
	*/
	memory := make(map[int]int)

	for _, num := range nums {
		// Create a map so that presence of element can be checked in O(1) time
		memory[num] = 1
	}

	result := 0

	for _, num := range nums {
		if memory[num-1] == 0 {
			// Element does not have a left neighbor,
			// which means it is the start of a sequence.
			sequenceLen := 1

			// Calculate length of sequence starting with num
			for i := num + 1; ; i++ {
				if memory[i] == 0 {
					sequenceLen++
				} else {
					break
				}
			}
			if sequenceLen > result {
				result = sequenceLen
			}
			// Avoid checking for num again if it is repeated in the array
			// This will help speed up the solution
			memory[num-1] = 2
		}
	}

	return result
}
