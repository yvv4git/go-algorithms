package main

import "fmt"

func uniquePaths(m int, n int) int {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(m*n), где m и n - размеры сетки, так как мы проходим по всей сетке ровно один раз.
		SPACE COMPLEXITY: O(m*n), так как мы создаем дополнительный массив dp размером m*n.
	*/
	// Создаем двумерный массив dp размером m*n
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Инициализируем первую клетку сетки
	dp[0][0] = 1

	// Заполняем первую строку сетки
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1]
	}

	// Заполняем первый столбец сетки
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0]
	}

	// Заполняем остальную часть сетки
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	// Возвращаем количество уникальных путей до правого нижнего угла
	return dp[m-1][n-1]
}

func main() {
	m := 3
	n := 7
	fmt.Println(uniquePaths(m, n)) // Выводит: 28
}
