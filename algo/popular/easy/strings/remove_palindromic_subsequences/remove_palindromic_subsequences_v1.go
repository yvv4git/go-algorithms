package remove_palindromic_subsequences

func removePalindromeSubV1(s string) int {
	/*
		METHOD: Iterate
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)
	*/
	if len(s) == 0 {
		return 0
	}

	for i := 0; i < len(s)/2; i++ {
		//fmt.Printf("[%s] [%d]->%c [%d]->%c \n", s, i, s[i], len(s)-1-i, s[len(s)-1-i])
		if s[i] != s[len(s)-1-i] {
			return 2
		}
	}

	return 1
}
