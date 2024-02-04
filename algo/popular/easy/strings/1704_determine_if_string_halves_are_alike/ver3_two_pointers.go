package main

func halvesAreAlikeV3(s string) bool {
	/*
		METHOD: Two pointers
		TIME COMPLEXITY: O(n), где n - длина строки, потому что мы проходим по каждому символу строки ровно один раз.
		Space complexity: O(1), потому что мы используем фиксированное количество дополнительной памяти, независимо от размера входных данных.
	*/
	n := len(s)
	l, r := 0, 0

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		l += isVowel(s[i])
		r += isVowel(s[j])
	}

	return l == r
}

func isVowel(x byte) int {
	if 'a' == x || 'e' == x || 'i' == x || 'o' == x || 'u' == x {
		return 1
	}
	if 'A' == x || 'E' == x || 'I' == x || 'O' == x || 'U' == x {
		return 1
	}
	return 0
}
