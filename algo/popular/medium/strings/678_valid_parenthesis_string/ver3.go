package _78_valid_parenthesis_string

func checkValidStringV3(s string) bool {
	/*
		METHOD: Loop with counters.
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	var mainStack int
	var helpStack int

	for _, c := range s {
		if c == '(' {
			mainStack++
		} else if c == '*' {
			if mainStack > 0 {
				mainStack--
				helpStack += 2
			} else {
				helpStack++
			}
		} else {
			if mainStack > 0 {
				mainStack--
			} else if helpStack > 0 {
				helpStack--
			} else {
				return false
			}
		}
	}

	return mainStack == 0
}
