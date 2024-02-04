package longest_nice_substring

func longestNiceSubstringV1(s string) string {
	/*
		METHOD: Bruteforce
		TIME COMPLEXITY: O(n^2)
		Space complexity: O(1)
	*/
	var lower, upper, longest, ind int

	for i := 0; i < len(s); i++ {
		lower, upper = 0, 0
		for j := i; j < len(s); j++ {
			if isLower(s[j]) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}

			if lower^upper == 0 {
				tmp := j - i + 1
				if tmp > longest {
					longest = tmp
					ind = i
				}
			}
		}
	}

	return s[ind : ind+longest]
}

func isLower(b byte) bool {
	// it's ok, since s consists of uppercase and lowercase English letters.
	return b >= 'a'
}
