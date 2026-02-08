package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(N, M int, grid [][]int) int {
	/*
		METHOD: Dynamic Programming (двумерная таблица)
		Суть задачи: Игрок движется по таблице N×M, может перемещаться только вправо или вниз.
		При прохождении через клетку снимается вес еды, записанный в этой клетке.
		Нужно найти минимальный суммарный вес для достижения правого нижнего угла.

		dp[i][j] - минимальный вес для достижения клетки (i, j).
		Переход: dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1]).
		Базовые случаи: dp[0][0] = grid[0][0], первая строка/столбец заполняются последовательно.

		TIME COMPLEXITY: O(n*m) - проход по всем ячейкам таблицы
		SPACE COMPLEXITY: O(n*m) - двумерный массив dp (можно оптимизировать до O(m))
	*/

	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, M)
	}

	// Заполняем первую ячейку
	dp[0][0] = grid[0][0]

	// Заполняем первую строку (можно двигаться только слева)
	for j := 1; j < M; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	// Заполняем первый столбец (можно двигаться только сверху)
	for i := 1; i < N; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// Заполняем остальные ячейки
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[N-1][M-1]
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var N, M int
	fmt.Fscan(in, &N, &M)

	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, M)
		for j := 0; j < M; j++ {
			fmt.Fscan(in, &grid[i][j])
		}
	}

	fmt.Println(solve(N, M, grid))
}
