package ver3_kmp

import "strings"

func shortestPalindrome(s string) string {
	/*
		METHOD: KMP(Knuth-Morris-Pratt)
		Time complexity: O(m + n), где m - длина подстроки, n - длина строки
		Space complexity: O(m)

		Time complexity
		Временная сложность O(m + n) в алгоритме KMP обусловлена двумя факторами:
		1. Построение префикс-функции
		Этот шаг требует O(m) времени, где m - длина подстроки, для которой строится префикс-функция.
		2. Поиск подстроки
		Этот шаг также требует O(n) времени, где n - длина строки, в которой ищется подстрока.

		Это связано с тем, что алгоритм KMP использует префикс-функцию для пропуска несовпадающих символов, что позволяет ему работать за линейное время.

		Space complexity
		Алгоритм KMP может использовать дополнительную память для хранения результатов, но это не влияет на общее пространственное потребление.

		KMP (Knuth-Morris-Pratt) алгоритм - это алгоритм поиска подстроки в строке.
		Он основан на принципе, что если некоторая часть строки не совпадает с подстрокой, то нет смысла проверять эту часть строки дальше.
		KMP алгоритм использует префикс-функцию, которая для каждого индекса в строке вычисляет максимальную длину префикса,
		который также является суффиксом для подстроки, начинающейся с этого индекса.
		Эта информация используется для перехода к следующему символу в строке, когда некоторая часть подстроки не совпадает с частью строки.
	*/
	sr := reverse(s)

	idx := getNext(s + "#" + sr)
	ans := sr + s[idx:]
	return ans
}

func reverse(s string) string {
	rev := strings.Builder{}
	for i := len(s) - 1; i >= 0; i-- {
		rev.WriteByte(s[i])
	}
	return rev.String()
}

func getNext(s string) int {
	n := len(s)

	next := make([]int, n)

	for i, j := 1, 0; i < n; i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}

		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next[n-1]
}
