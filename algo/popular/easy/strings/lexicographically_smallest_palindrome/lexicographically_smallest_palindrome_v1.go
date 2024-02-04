package lexicographically_smallest_palindrome

func makeSmallestPalindromeV1(s string) string {
	/*
		METHOD: Swap
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	res := []byte(s)

	for i := 0; i < len(s)/2; i++ {
		if a, b := res[i], res[len(s)-1-i]; a != b {
			if a < b {
				res[len(s)-1-i] = a
			} else {
				res[i] = b
			}
		}
	}

	return string(res)
}
