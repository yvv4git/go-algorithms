package _42_valid_anagram

func isAnagramV3(s string, t string) bool {
	/*
		METHOD: Hashmap
		Time complexity : O(n)
		Space complexity : O(n)
	*/
	sr := map[byte]int{}
	tr := map[byte]int{}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		sr[s[i]] += 1
		tr[t[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		if sr[s[i]] != tr[s[i]] {
			return false
		}
	}

	return true
}
