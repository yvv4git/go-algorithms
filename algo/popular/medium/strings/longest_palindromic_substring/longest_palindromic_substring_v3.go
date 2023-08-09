package longest_palindromic_substring

func longestPalindromeV3(s string) string {
	/*
		Method: Algorithm Manacher's
		Time complexity: O(n)
		Space complexity: O(n)
	*/
	n := len(s)
	d1 := make([]int, n) //odd
	d2 := make([]int, n) //even
	d1[0] = 1
	d2[0] = 1
	x1 := 0
	x2 := 0
	maxLength1 := 1
	idx1 := 0
	maxLength2 := 0
	idx2 := 0
	for i := 1; i < n; i++ {
		var ln1 int
		var ln2 int
		// odd
		if x1+d1[x1]-1 < i {
			ln1 = 1
		} else {
			ln1 = min(d1[2*x1-i], x1+d1[x1]-i)
		}

		for i-ln1 >= 0 && i+ln1 < n && s[i-ln1] == s[i+ln1] {
			ln1++
		}
		d1[i] = ln1
		if i+d1[i] > x1+d1[x1] {
			x1 = i
		}
		if d1[i] > maxLength1 {
			maxLength1 = d1[i]
			idx1 = i
		}
		//even
		if s[i] != s[i-1] {
			continue
		}

		if x2+d2[x2]-1 < i {
			ln2 = 1
		} else {
			ln2 = min(d2[2*x2-i], x2+d2[x2]-i)
		}
		for i-ln2-1 >= 0 && i+ln2 < n && s[i-ln2-1] == s[i+ln2] {
			ln2++
		}
		d2[i] = ln2
		if i+d2[i] > x2+d2[x2] {
			x2 = i
		}
		if d2[i] > maxLength2 {
			maxLength2 = d2[i]
			idx2 = i
		}
	}
	ans := make([]rune, 0)
	if maxLength1 > maxLength2 {
		for i := idx1 - maxLength1 + 1; i < idx1+maxLength1; i++ {
			ans = append(ans, rune(s[i]))
		}
	} else {
		for i := idx2 - maxLength2; i < idx2+maxLength2; i++ {
			ans = append(ans, rune(s[i]))
		}
	}
	return string(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
