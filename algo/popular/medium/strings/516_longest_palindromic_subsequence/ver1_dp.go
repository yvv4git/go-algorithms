package _16_longest_palindromic_subsequence

func longestPalindromeSubseqV1(s string) int {
	/*
		METHOD: Dynamic programming
		Time complexity: O(n^2)
		Space complexity: O(n^2)

		Time complexity
		Временная сложность этого алгоритма - O(n^2), где n - длина строки s.
		Это связано с двумя вложенными циклами, которые проходят по всем символам строки.

		Space complexity
		Пространственная сложность - O(n^2), так как мы используем двумерный массив dp для хранения результатов подсчета.
	*/
	n := len(s) // Вычисляем длину строки s

	// Создаем двумерный массив dp размером n x n, заполненный нулями
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Заполняем таблицу dp
	// Для каждого символа в строке s
	for i := n - 1; i >= 0; i-- {
		// Сам на себя символ является палиндромом, поэтому dp[i][i] = 1
		dp[i][i] = 1

		// Для каждого символа, следующего за i-м
		for j := i + 1; j < n; j++ {
			// Если символы i и j равны, то dp[i][j] = 2 + dp[i+1][j-1]
			// (2 - это символы i и j, dp[i+1][j-1] - наибольшая палиндромная подпоследовательность в строке между i и j)
			if s[i] == s[j] {
				dp[i][j] = 2 + dp[i+1][j-1]
			} else {
				// Если символы i и j не равны, то dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				// (max - это наибольшая палиндромная подпоследовательность в строке между i и j)
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	// Возвращаем наибольшую палиндромную подпоследовательность в строке s
	return dp[0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
