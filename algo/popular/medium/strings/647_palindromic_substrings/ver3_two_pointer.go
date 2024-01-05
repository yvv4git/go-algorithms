package _47_palindromic_substrings

func countSubstringsV3(s string) int {
	/*
		Method: Two pointer
		Time complexity: O(n^2)
		Space complexity: O(1)

		Time complexity
		Временная сложность этого алгоритма составляет O(n^2), где n - длина строки.
		Это связано с двумя циклами: внешним, который проходит по всем символам строки, и внутренним,
		который проходит по всем индексам, соответствующим одному символу.

		Space complexity
		Пространственная сложность составляет O(1), так как мы используем только несколько переменных для хранения информации о палиндромах.
		Мы не используем дополнительное пространство, которое зависит от размера входных данных.
	*/
	return solution{s}.countSubstrings()
}

type solution struct {
	s string
}

func (s solution) countSubstrings() int {
	count := 0
	for i := 0; i < len(s.s); i++ {
		count += s.findPalindrom(i, i)
		count += s.findPalindrom(i, i+1)
	}
	return count
}

func (s solution) isOverflow(l, r int) bool {
	return l < 0 || r >= len(s.s)
}

func (s solution) findPalindrom(l, r int) int {
	count := 0
	maxLength := 0
	for !s.isOverflow(l, r) && s.s[l] == s.s[r] {
		if r-l+1 > maxLength {
			count++
			maxLength = r - l + 1
		}
		l--
		r++
	}
	return count
}
