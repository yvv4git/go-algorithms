package _78_valid_parenthesis_string

func checkValidStringV2(s string) bool {
	/*
		Method: Loop with counters.
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	open, closed, last := 0, 0, len(s)-1

	for i, c := range s {
		if c == '(' || c == '*' {
			open++
		} else {
			open--
		}

		if s[last-i] == ')' || s[last-i] == '*' {
			closed++
		} else {
			closed--
		}

		if closed < 0 || open < 0 {
			return false
		}
	}

	return true
}
