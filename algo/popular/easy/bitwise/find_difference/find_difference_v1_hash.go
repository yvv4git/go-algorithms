package find_difference

func findTheDifferenceV1(s string, t string) byte {
	/*
		Method: Hash
		Time complexity: O(n + n + n) = O(3n) = O(n)
		Space complexity: O(n)
	*/

	// Space: O(n + n) = O(2n) = O(n)
	ms, mt := make(map[byte]int), make(map[byte]int)
	// Time: O(n)
	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}
	// Time: O(n)
	for i := 0; i < len(t); i++ {
		mt[t[i]]++
	}
	// Time: O(n)
	for k, v := range mt {
		// Time: O(1)
		if c, ok := ms[k]; !ok || v != c {
			return k
		}
	}

	var b byte
	return b
}
