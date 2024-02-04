package longest_palindrome

func longestPalindromeV1(s string) int {
	/*
		METHOD: Map
		TIME COMPLEXITY: O(n), n - length of string
		SPACE COMPLEXITY: O(n), if all characters are different in string
	*/
	result := 0
	var m map[rune]int
	m = make(map[rune]int)

	for _, ch := range s {
		if c, ok := m[ch]; !ok {
			m[ch] = 1
		} else {
			m[ch] = c + 1
		}
	}

	odd := false
	for _, v := range m {
		if v%2 == 0 {
			result += v
		} else {
			result += v - 1
			odd = true
		}
	}

	if odd {
		return result + 1
	}

	return result
}
