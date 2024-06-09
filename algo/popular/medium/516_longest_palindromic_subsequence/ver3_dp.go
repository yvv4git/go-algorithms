package _16_longest_palindromic_subsequence

// Функция longestPalindromeSubseqV3 находит длину наибольшей палиндромной подпоследовательности в строке s
func longestPalindromeSubseqV3(s string) int {
	/*
		METHOD: Dynamic programming + Recursion
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(n^2)

		Time complexity
		Временная сложность этого алгоритма - O(n^2), где n - длина строки s.
		Это связано с двумя вложенными циклами, которые проходят по всем символам строки.

		Space complexity
		Пространственная сложность - O(n^2), так как мы используем двумерный массив dp для хранения результатов подсчета.
	*/

	// Если длина строки s меньше 2, то возвращаем длину строки s
	if len(s) < 2 {
		return len(s)
	}

	// Создаем двумерный массив dp размером len(s) x len(s)
	dp := make([][]int, len(s))

	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	// Вызываем функцию helper, которая вычисляет длину наибольшей палиндромной подпоследовательности
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
