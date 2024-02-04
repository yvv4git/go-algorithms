package _78_valid_parenthesis_string

func checkValidStringV4(s string) bool {
	/*
		METHOD: Greedy solution
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	leftMin, leftMax := 0, 0

	for _, symbol := range s {
		switch symbol {
		case '(':
			leftMin, leftMax = leftMin+1, leftMax+1
		case ')':
			leftMin, leftMax = leftMin-1, leftMax-1
		case '*':
			leftMin, leftMax = leftMin-1, leftMax+1
		}

		if leftMin < 0 {
			leftMin = 0
		}

		if leftMax < 0 {
			return false
		}
	}

	return leftMin == 0
}
