package find_difference

func findTheDifferenceV2(s string, t string) byte {
	/*
		Method: Bitwise
		Time complexity: O(n + n) = O(2n) = O(n)
		Space complexity: O(n)
	*/
	var bs, bt byte
	// Time: O(n)
	for i := 0; i < len(s); i++ {
		bs += s[i]
	}

	// Time: O(n)
	for i := 0; i < len(t); i++ {
		bt += t[i]
	}

	return bt - bs
}