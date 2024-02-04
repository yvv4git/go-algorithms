package __longest_palindromic_substring

func longestPalindromeV2(s string) string {
	/*
		METHOD: Hash
		TIME COMPLEXITY: O(n^2)
		Space complexity: O(n)
	*/
	var n = len(s)

	var matrix = make([][]bool, n)
	var result string

	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)

		matrix[i][i] = true
		result = string(s[i])
	}
	var longestPalindromeLength = 0

	for start := n - 1; start >= 0; start-- {
		for end := start + 1; end < n; end++ {
			if string(s[start]) == string(s[end]) {
				if end-start == 1 || matrix[start+1][end-1] {
					matrix[start][end] = true
					if longestPalindromeLength < end-start+1 {
						longestPalindromeLength = end - start + 1
						result = string(s[start : end+1])
					}
				}
			}
		}
	}

	return result
}
