package longest_palindrome

func longestPalindromeV2(s string) int {
	/*
		Method: Use bool array
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	hash := make([]bool, 58) // O(1), because we have constant length
	result := 0

	for _, char := range s {
		if hash[int(char-'A')] {
			hash[int(char-'A')] = false
			result += 2
			continue
		}
		hash[int(char-'A')] = true
	}

	if len(s) > result {
		result++
	}

	return result
}
