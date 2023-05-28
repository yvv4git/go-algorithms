package decode_xored_array_v1

func decodeV1(encoded []int, first int) []int {
	/*
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	result := make([]int, len(encoded)+1)
	result[0] = first

	for idx, summa := range encoded {
		result[idx+1] = summa ^ result[idx]
	}

	return result
}
