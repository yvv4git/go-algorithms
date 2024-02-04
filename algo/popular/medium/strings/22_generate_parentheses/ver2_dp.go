package main

// Функция для генерации всех возможных комбинаций скобок.
func generateParenthesisV2(n int) []string {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(4^n / n^(3/2))
		Space complexity: O(4^n / n^(3/2))

		Time complexity
		Временная сложность O(4^n / n^(3/2)) обусловлена тем, что для каждой пары скобок n,
		мы генерируем все возможные комбинации скобок, которые могут быть сформированы из n пар скобок.
		Количество таких комбинаций равно количеству Каталана, которое приблизительно равно 4^n / n^(3/2).

		Space complexity
		Пространственная сложность O(4^n / n^(3/2)) обусловлена тем, что мы должны хранить все сгенерированные комбинации скобок.
		Количество комбинаций примерно равно 4^n / n^(3/2), поэтому пространственная сложность соответствует этому количеству.
	*/
	// База динамического программирования: dp[i] будет содержать все комбинации скобок для i пар скобок.
	dp := make([][]string, n+1)
	dp[0] = []string{""}   // 0 пар скобок - это одна пустая строка
	dp[1] = []string{"()"} // 1 пара скобок - это одна скобка

	// Заполняем dp для каждого i от 2 до n.
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			// Для каждого j мы можем разместить i-j-1 пару скобок внутри i-j-1 пару скобок и j пар скобок за пределами.
			for _, inside := range dp[j] {
				for _, outside := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+inside+")"+outside)
				}
			}
		}
	}

	return dp[n]
}
