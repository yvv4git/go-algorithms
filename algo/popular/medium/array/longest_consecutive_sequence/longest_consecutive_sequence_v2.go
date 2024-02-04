package longest_consecutive_sequence

func longestConsecutiveV2(nums []int) int {
	/*
		TIME COMPLEXITY: O(n). It consists of building a memory in O(n) time + iteration in O(n).
		Space complexity: at worst nnn elements will be stored in the memory, so O(n).
	*/
	memory := make(map[int]int)
	result := 0

	for i := 0; i < len(nums); i++ {
		memory[nums[i]] = nums[i]
	}

	for key, _ := range memory {
		_, ok := memory[key-1]
		if !ok {
			current := key
			count := 1

			for {
				_, not := memory[current+1]
				if !not {
					break
				}

				current++
				count++
			}

			if result < count {
				result = count
			}
		}
	}

	return result
}
