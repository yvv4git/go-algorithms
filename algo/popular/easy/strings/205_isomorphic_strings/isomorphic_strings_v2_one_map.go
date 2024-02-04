package _05_isomorphic_strings

func isIsomorphicV2(s string, t string) bool {
	/*
		METHOD: Map / Hash
		TIME COMPLEXITY: O(2n)
		Space complexity: O(1), количество символов ограничено
	*/
	return helper(s, t) && helper(t, s)
}

func helper(s string, t string) bool {
	charMap := make(map[byte]byte)
	for i := range s {
		if v, ok := charMap[s[i]]; !ok {
			charMap[s[i]] = t[i]
		} else {
			if v != t[i] {
				return false
			}
		}
	}

	return true
}
