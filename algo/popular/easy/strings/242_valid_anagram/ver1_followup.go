package _42_valid_anagram

func isAnagramV1(s string, t string) bool {
	/*
		METHOD: Followup
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		chars[v]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
