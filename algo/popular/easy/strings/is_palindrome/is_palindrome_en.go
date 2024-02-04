package is_palindrome

func isPalindromeEn(input string) (out bool) {
	/*
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	l := 0
	r := len(input) - 1

	for l < r {
		if input[l] != input[r] {
			return false
		}

		l += 1
		r -= 1
	}

	return true
}
