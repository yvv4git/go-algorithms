package longest_consecutive_sequence

func longestConsecutiveV1(nums []int) int {
	/*
		TIME COMPLEXITY: O(n). It consists of building a set in O(n) time + iteration in O(n).
		Space complexity: at worst nnn elements will be stored in the set, so O(n).
	*/

	// Construct a set out of the nums array.
	numsSet := make(map[int]struct{})
	for _, n := range nums {
		numsSet[n] = struct{}{}
	}

	// The answer is stored here.
	maxSequenceLen := 0

	// Iterate through the set.
	for n := range numsSet {
		// We check if n-1 is in the set. If it is, then n is not the beginning of a sequence
		// and we go to the next number immediately.
		if _, ok := numsSet[n-1]; !ok {
			// Otherwise, we increment n in a loop to see if the next consecutive value is stored in nums.
			seqLen := 1
			for {
				if _, ok = numsSet[n+seqLen]; ok {
					seqLen++
					continue
				}
				// When the sequence is over, see if we did better than before.
				maxSequenceLen = max(seqLen, maxSequenceLen)
				break
			}
		}
	}

	return maxSequenceLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
