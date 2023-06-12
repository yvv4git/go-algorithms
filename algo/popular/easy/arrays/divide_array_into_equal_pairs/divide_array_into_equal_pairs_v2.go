package divide_array_into_equal_pairs

func divideArrayV2(nums []int) bool {
	/*
		Method: Hash
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	hash := map[int]bool{}

	for _, n := range nums {
		if hash[n] {
			delete(hash, n)
		} else {
			hash[n] = true
		}
	}

	return len(hash) == 0
}
