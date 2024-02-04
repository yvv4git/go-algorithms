package decode_xored_array

func decodeV2(encoded []int, first int) []int {
	/*
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	origin := []int{first}

	for _, code := range encoded {
		source := origin[len(origin)-1] ^ code
		origin = append(origin, source)
	}

	return origin
}
