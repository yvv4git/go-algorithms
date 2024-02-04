package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(m * n), где m и n - длины строк s1 и s2 соответственно.
		Space complexity: O(m * n), так как мы используем дополнительный двумерный массив dp размером (m + 1) * (n + 1).
	*/
	// Проверяем, что суммарная длина s1 и s2 равна длине s3
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// Создаем двумерный массив dp
	dp := make([][]bool, len(s1)+1)
	for i := range dp {
		dp[i] = make([]bool, len(s2)+1)
	}

	// Базовый случай: пустая строка всегда взаимна с любой другой строкой
	dp[0][0] = true

	// Заполняем первую строку dp
	for i := 1; i <= len(s1); i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}

	// Заполняем первый столбец dp
	for j := 1; j <= len(s2); j++ {
		if s2[j-1] == s3[j-1] {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Заполняем остальную часть dp
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			// Если текущий символ s3 совпадает с символом s1 и s2,
			// то мы можем получить текущую подстроку s3 из объединения
			// подстрок s1 и s2 без текущего символа или из подстроки s1
			// без текущего символа или из подстроки s2 без текущего символа
			dp[i][j] = (s3[i+j-1] == s1[i-1] && dp[i-1][j]) ||
				(s3[i+j-1] == s2[j-1] && dp[i][j-1])
		}
	}

	// Возвращаем результат для всей строки s3
	return dp[len(s1)][len(s2)]
}
