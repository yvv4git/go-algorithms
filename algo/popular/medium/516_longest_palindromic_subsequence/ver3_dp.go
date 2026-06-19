package _16_longest_palindromic_subsequence

func longestPalindromeSubseqV3(s string) int {
	/*
		TASK: Найти длину наибольшей палиндромной подпоследовательности в строке s.
		Подпоследовательность отличается от подстроки тем, что символы не обязаны
		идти подряд (можно пропускать символы).

		APPROACH: Top-down DP (Recursion + Memoization, с указателем на матрицу)
		Вариант ver2, но передаём dp через указатель (*[][]int) вместо замыкания.
		1. Рекурсивная функция helper(dp, s, begin, end).
		2. База: begin > end — 0, begin == end — 1.
		3. Если s[begin] == s[end]: dp[begin][end] = 2 + helper(begin+1, end-1).
		4. Иначе: dp[begin][end] = max(helper(begin+1, end), helper(begin, end-1)).
		5. Ответ: helper(&dp, s, 0, n-1).

		TIME COMPLEXITY: O(n^2)
		- Каждая пара (begin, end) вычисляется один раз.

		SPACE COMPLEXITY: O(n^2)
		- Таблица dp размером n×n плюс глубина рекурсии O(n).
	*/

	if len(s) < 2 {
		return len(s)
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	return helper(&dp, s, 0, len(s)-1)
}

// Функция helper вычисляет длину наибольшей палиндромной подпоследовательности в строке s
func helper(dp *[][]int, s string, begin int, end int) int {
	// Если индекс begin больше end, то возвращаем 0
	if begin > end {
		return 0
	}

	// Если индексы begin и end равны, то возвращаем 1
	if begin == end {
		return 1
	}

	// Если значение в dp[begin][end] равно 0, то вычисляем его
	if (*dp)[begin][end] == 0 {
		// Если символы begin и end равны, то dp[begin][end] = 2 + helper(dp, s, begin+1, end-1)
		if s[begin] == s[end] {
			(*dp)[begin][end] = 2 + helper(dp, s, begin+1, end-1)
		} else {
			// Если символы begin и end не равны, то dp[begin][end] = max(helper(dp, s, begin+1, end), helper(dp, s, begin, end-1))
			(*dp)[begin][end] = maxV3(helper(dp, s, begin+1, end), helper(dp, s, begin, end-1))
		}
	}

	// Возвращаем значение в dp[begin][end]
	return (*dp)[begin][end]
}

// Функция maxV3 возвращает максимальное из двух чисел
func maxV3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
