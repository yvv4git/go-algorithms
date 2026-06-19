package _16_longest_palindromic_subsequence

func longestPalindromeSubseqV1(s string) int {
	/*
		TASK: Найти длину наибольшей палиндромной подпоследовательности в строке s.
		Подпоследовательность отличается от подстроки тем, что символы не обязаны
		идти подряд (можно пропускать символы).

		APPROACH: Dynamic Programming (Bottom-up, 2D DP)
		1. Создаем таблицу dp[i][j] — длина наибольшей палиндромной подпоследовательности
		   в подстроке s[i..j] (i — начало, j — конец).
		2. База: dp[i][i] = 1 (один символ — палиндром длины 1).
		3. Заполняем таблицу снизу вверх (i от n-1 до 0, j от i+1 до n-1):
		   - Если s[i] == s[j]: dp[i][j] = 2 + dp[i+1][j-1]
		   - Иначе: dp[i][j] = max(dp[i+1][j], dp[i][j-1])
		4. Ответ в dp[0][n-1].

		TIME COMPLEXITY: O(n^2)
		- Два вложенных цикла по всем n символам.

		SPACE COMPLEXITY: O(n^2)
		- Таблица dp размером n×n.
	*/
	n := len(s)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1

		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
