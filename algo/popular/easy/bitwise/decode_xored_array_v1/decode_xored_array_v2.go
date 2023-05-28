package decode_xored_array_v1

func decodeV2(encoded []int, first int) []int {
	/*
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	origin := []int{first}

	for _, code := range encoded {
		source := origin[len(origin)-1] ^ code
		origin = append(origin, source)
	}

	return origin
}
