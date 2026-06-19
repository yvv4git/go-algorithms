package _16_longest_palindromic_subsequence

func longestPalindromeSubseqV2(s string) int {
	/*
		TASK: Найти длину наибольшей палиндромной подпоследовательности в строке s.
		Подпоследовательность отличается от подстроки тем, что символы не обязаны
		идти подряд (можно пропускать символы).

		APPROACH: Top-down DP (Recursion + Memoization)
		1. Используем рекурсивную функцию compute(s, i, j, DP), которая возвращает
		   длину наибольшей палиндромной подпоследовательности в s[i..j].
		2. База: если i == j — 1, если j < i — 0.
		3. Если s[i] == s[j]: результат = 2 + compute(i+1, j-1).
		4. Иначе: результат = max(compute(i+1, j), compute(i, j-1)).
		5. Кешируем результаты в DP[i][j] (изначально -1).
		6. Ответ: compute(s, 0, n-1, DP).

		TIME COMPLEXITY: O(n^2)
		- Каждая пара (i, j) вычисляется один раз.

		SPACE COMPLEXITY: O(n^2)
		- Таблица DP размером n×n плюс глубина рекурсии O(n).
	*/

	DP := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		DP[i] = make([]int, len(s))
		for j := 0; j < len(s); j++ {
			DP[i][j] = -1
		}
	}
	return compute(s, 0, len(s)-1, DP)
}

func compute(s string, i, j int, DP [][]int) int {
	if j == i {
		return 1
	}
	if j < i {
		return 0
	}
	if DP[i][j] != -1 {
		return DP[i][j]
	}
	if s[i] == s[j] {
		DP[i][j] = 2 + compute(s, i+1, j-1, DP)
	} else {
		DP[i][j] = max(compute(s, i+1, j, DP), compute(s, i, j-1, DP))
	}
	return DP[i][j]
}


