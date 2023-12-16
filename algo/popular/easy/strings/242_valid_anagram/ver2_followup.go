package _42_valid_anagram

func isAnagramV2(s string, t string) bool {
	/*
		Method: Followup
		Time complexity : O(n)
		Space complexity : O(1)
	*/
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}

	for _, v := range t {
		if number, exists := chars[v]; !exists || number == 0 {
			return false
		}
		chars[v]--
	}

	return len(s) == len(t)
}
