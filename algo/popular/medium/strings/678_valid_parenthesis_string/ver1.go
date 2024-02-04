package _78_valid_parenthesis_string

func checkValidStringV1(s string) bool {
	/*
		METHOD: Reverse and check.
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	var validate func(string, byte) bool
	validate = func(s string, op byte) bool {
		nOpen, nClose, nFree := 0, 0, 0

		for i := 0; i < len(s); i++ {
			if s[i] == op {
				nOpen++
			} else if s[i] == '*' {
				nFree++
			} else {
				nClose++
			}

			if nOpen+nFree < nClose {
				return false
			}
		}

		return true
	}

	var reverse func(string) string
	reverse = func(s string) string {
		if len(s) == 0 {
			return s
		}

		// SPACE COMPLEXITY: O(n)
		origin := []byte(s)

		l, r := 0, len(s)-1

		for l < r {
			origin[l], origin[r] = origin[r], origin[l]
			l++
			r--
		}

		return string(origin)
	}

	return validate(s, '(') && validate(reverse(s), ')')
}
