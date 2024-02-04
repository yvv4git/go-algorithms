package lexicographically_smallest_palindrome

func makeSmallestPalindromeV2(s string) string {
	/*
		METHOD: Swap
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	p := []byte(s)

	for i, j := 0, len(p)-1; i < j; i, j = i+1, j-1 {
		if p[i] > p[j] {
			p[i] = p[j]
		} else {
			p[j] = p[i]
		}
	}

	return string(p)
}
