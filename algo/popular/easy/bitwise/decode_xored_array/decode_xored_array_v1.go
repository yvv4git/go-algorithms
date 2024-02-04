package decode_xored_array

func decodeV1(encoded []int, first int) []int {
	/*
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	result := make([]int, len(encoded)+1)
	result[0] = first

	for idx, summa := range encoded {
		result[idx+1] = summa ^ result[idx]
	}

	return result
}
