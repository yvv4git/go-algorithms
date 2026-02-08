package main

import (
	"fmt"
)

func solve(N, M int) int {
	/*
		METHOD: Dynamic Programming (двумерная таблица)
		Суть задачи: Конь на доске N×M может двигаться только: (↓2, →1) или (↓1, →2).
		Нужно найти количество маршрутов из левого верхнего в правый нижний угол.

		dp[i][j] - количество способов добраться до клетки (i, j).
		Переход: dp[i][j] = dp[i-2][j-1] + dp[i-1][j-2] (если эти клетки существуют).

		TIME COMPLEXITY: O(n*m) - проход по всем ячейкам таблицы
		SPACE COMPLEXITY: O(n*m) - двумерный массив dp
	*/

	if N == 1 && M == 1 {
		return 1
	}
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, M+1)
	}
	dp[1][1] = 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if i == 1 && j == 1 {
				continue
			}
			// from up-left (two down, one right)
			if i-2 >= 1 && j-1 >= 1 {
				dp[i][j] += dp[i-2][j-1]
			}
			// from up-left (one down, two right)
			if i-1 >= 1 && j-2 >= 1 {
				dp[i][j] += dp[i-1][j-2]
			}
		}
	}
	return dp[N][M]
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	fmt.Println(solve(N, M))
}
