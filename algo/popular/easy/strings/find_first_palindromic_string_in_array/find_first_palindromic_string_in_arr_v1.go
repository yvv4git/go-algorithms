package find_first_palindromic_string_in_array

func firstPalindromeV1(words []string) string {
	/*
		METHOD: Iterate
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)
	*/
	// Проверяем строку на палиндром
	isPalindromic := func(s string) bool {
		if len(s) < 2 {
			return true
		}

		l, r := 0, len(s)-1

		for l < r {
			if s[l] != s[r] {
				return false
			}

			l, r = l+1, r-1
		}

		return true
	}

	for _, word := range words {
		if isPalindromic(word) {
			return word
		}
	}

	return ""
}
