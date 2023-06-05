package find_difference

import "strings"

func findTheDifferenceV3(s string, t string) byte {
	/*
		Method: Bitwise
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	// Concatenate two string
	var str strings.Builder
	str.WriteString(s)
	str.WriteString(t)

	var resultByte int
	// Time complexity: O(n)
	for _, value := range str.String() {
		resultByte ^= int(value)
	}

	return byte(resultByte)
}
