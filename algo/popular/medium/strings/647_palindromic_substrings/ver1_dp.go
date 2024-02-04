package _47_palindromic_substrings

func countSubstringsV1(s string) int {
	/*
		METHOD: Dynamic programming
		Time complexity: O(n^2)
		Space complexity: O(n^2)

		Time complexity
		Временная сложность этого алгоритма - O(n^2), где n - длина строки.
		Это связано с двумя вложенными циклами, которые проходят по всем символам строки.

		Space complexity
		Пространственная сложность - O(n^2), так как мы используем двумерный массив dp для хранения результатов.

		Метод динамического программирования используется здесь, потому что мы можем использовать результаты предыдущих вычислений для вычисления текущих.
		Например, мы можем использовать результаты для подстроки s[i+1:j-1] для определения, является ли подстрока s[i:j] палиндромом.
	*/
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	count := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			// Если длина подстроки меньше или равна 2, то она является палиндромом, если символы равны
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				count++
			}
		}
	}

	return count
}
