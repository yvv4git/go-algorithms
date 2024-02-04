package main

import "strconv"

func isStrictlyPalindromicV2(n int) bool {
	/*
		METHOD: Iterate
		TIME COMPLEXITY: O(n^2), так как цикла в цикле
		Space complexity: O(1)
	*/
	for i := 2; i <= n-2; i++ {
		if !isPallindrome(n, i) {
			return false
		}
	}
	return true
}

func isPallindrome(n, i int) bool {
	var x string

	for n > 0 {
		x += strconv.Itoa(n % i)
		n /= i
	}

	s := 0
	e := len(x) - 1

	for s < e {
		if x[s] != x[e] {
			return false
		}
		s++
		e--
	}

	return true
}
