package remove_palindromic_subsequences

func removePalindromeSubV3(s string) int {
	/*
		Method: Iterate
		Time complexity: O(n)
		Space complexity: O(1)

		1. If s is a palindrome return 1
		2. If not, remove all the letter a first, then remove all the letter b
		3. Thus, the maximum steps is 2 if s is not a palindrome
	*/
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return 2
		}
	}

	return 1
}
