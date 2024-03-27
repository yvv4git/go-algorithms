package top_k_frequently_elements

func topKFrequentV3(nums []int, k int) []int {

	// (1) count repetitions ans save them into map
	//where key is the number, and value is the count of repetitions
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// (2) convert map of counts into slice
	//where index is the count, and value is the slice of numbers with the same count of repetitions
	freqs := make([][]int, 1+len(nums))
	for key, val := range count {
		freqs[val] = append(freqs[val], key)
	}

	// (3) go backwards through the slice of frequences and add non-null values to result
	// we are going backwards as the index equals frequency, and we need top K most frequent numbers
	res := make([]int, 0, k)
	for i := len(freqs) - 1; i >= 0; i-- {
		if freqs[i] != nil {
			res = append(res, freqs[i]...)
			if len(res) >= k {
				res = res[:k]
				break
			}
		}
	}

	return res
}
