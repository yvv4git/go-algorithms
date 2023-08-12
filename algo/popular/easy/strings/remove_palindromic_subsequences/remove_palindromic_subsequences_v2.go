package remove_palindromic_subsequences

func removePalindromeSubV2(s string) int {
	/*
		Method: Iterate
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	if len(s) == 0 {
		return 0
	}

	for i, j := 0, len(s)-1; i < j; i++ {
		if s[i] != s[j] {
			return 2
		}
		j--
	}

	return 1
}
