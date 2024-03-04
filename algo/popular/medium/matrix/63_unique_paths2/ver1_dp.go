package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	/*
		METHOD: Dynamic programming
		В этом коде мы сначала инициализируем матрицу dp нулями, а затем заполняем ее динамическим программированием.
		Если клетка (i, j) не заблокирована, то количество уникальных путей до нее равно сумме путей до соседних клеток (i-1, j) и (i, j-1).
		В конце мы возвращаем количество уникальных путей до нижнего правого угла.

		TIME COMPLEXITY: O(m*n), где m и n - количество строк и столбцов в матрице obstacleGrid соответственно.
		Это связано с тем, что мы проходим по каждой клетке в матрице ровно один раз.

		SPACE COMPLEXITY: O(m*n), так как мы создаем матрицу dp размером m*n.
	*/
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	// Создаем матрицу dp и инициализируем ее нулями
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Если начальная клетка не заблокирована, то устанавливаем dp[0][0] в 1
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	}

	// Заполняем первую строку dp
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = dp[0][j-1]
		}
	}

	// Заполняем первый столбец dp
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	// Заполняем остальную часть матрицы dp
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	// Возвращаем количество уникальных путей до нижнего правого угла
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid)) // Выводит: 2
}
