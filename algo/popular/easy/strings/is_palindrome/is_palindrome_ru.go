package is_palindrome

func isPalindromeRu(s string) bool {
	var runes []rune
	for _, r := range s {
		runes = append(runes, r)
	}

	l, r := 0, len(runes)-1
	for l < r {
		if runes[l] != runes[r] {
			return false
		}
		l++
		r--
	}

	return true
}
