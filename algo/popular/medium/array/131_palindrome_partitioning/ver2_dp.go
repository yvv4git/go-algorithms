package _31_palindrome_partitioning

func partitionV2(s string) [][]string {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(n^2)

		Time complexity
		Временная сложность этого алгоритма - O(n^2), где n - длина строки s.
		Это связано с двумя циклами, которые проходят по всем символам строки.

		Space complexity
		Пространственная сложность - O(n^2), так как мы используем двумерный массив dp размером n x n.
	*/
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// Инициализируем dp для подстрок длины 1 и 2
	for i := 0; i < n; i++ {
		dp[i][i] = true
		if i < n-1 && s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	// Заполняем dp для подстрок длины больше 2
	for length := 3; length <= n; length++ {
		for i := 0; i+length-1 < n; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			}
		}
	}

	// Используем dp для разбиения строки на палиндромы
	var result [][]string
	var path []string
	dfs(s, 0, dp, &path, &result)
	return result
}

func dfs(s string, start int, dp [][]bool, path *[]string, result *[][]string) {
	if start == len(s) {
		*result = append(*result, append([]string{}, *path...))
		return
	}
	for i := start; i < len(s); i++ {
		if dp[start][i] {
			*path = append(*path, s[start:i+1])
			dfs(s, i+1, dp, path, result)
			*path = (*path)[:len(*path)-1]
		}
	}
}
