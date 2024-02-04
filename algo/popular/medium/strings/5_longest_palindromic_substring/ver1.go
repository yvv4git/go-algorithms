package __longest_palindromic_substring

func longestPalindromeV1(s string) string {
	/*
		METHOD:
		Time complexity: O(n^2)
		Space complexity:  O(1)
	*/
	var result string

	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; i <= j; j-- {
			if s[i] != s[j] {
				continue
			}

			if isPalindrome(i, j, s) {
				palindrome := s[i : j+1]

				if len(result) < len(palindrome) {
					result = palindrome
				}

				break
			}
		}
	}

	return result
}

func isPalindrome(start, end int, s string) bool {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
