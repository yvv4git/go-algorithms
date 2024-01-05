package _47_palindromic_substrings

func countSubstringsV2(s string) int {
	/*
		Method: Hashmap
		Time complexity: O(n^2)
		Space complexity: O(n)

		Time complexity
		Временная сложность этого алгоритма составляет O(n^2), где n - длина строки.
		Это связано с двумя циклами: внешним, который проходит по всем символам строки, и внутренним,
		который проходит по всем индексам, соответствующим одному символу.

		Space complexity
		Пространственная сложность составляет O(n), так как мы используем hashmap для хранения индексов всех символов в строке.
	*/
	var res int

	m := make(map[byte][]int)

	for i := 0; i < len(s); i++ {
		m[s[i]] = append(m[s[i]], i)

		if len(m[s[i]]) > 1 {
			for _, val := range m[s[i]] {
				if val != i && isPalindromic(s[val:i+1]) {
					res++
				}
			}
		}
	}

	return res + len(s)
}

func isPalindromic(s string) bool {
	if len(s) == 1 || len(s) == 0 {
		return true
	}

	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}
