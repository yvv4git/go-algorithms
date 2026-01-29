package main

import (
	"fmt"
)

func solve(N, M int) int64 {
	if N == 1 && M == 1 {
		return 1
	}

	// Создаем массив dp[строка][столбец]
	dp := make([][]int64, N+1)
	for i := range dp {
		dp[i] = make([]int64, M+1)
	}

	dp[1][1] = 1

	// visited[i][j] = true если клетка уже была полностью обработана
	visited := make([][]bool, N+1)
	for i := range visited {
		visited[i] = make([]bool, M+1)
	}

	// Обрабатываем пока все клетки не стабилизируются
	changed := true
	for changed {
		changed = false

		for j := 1; j <= M; j++ {
			for i := 1; i <= N; i++ {
				if visited[i][j] || dp[i][j] == 0 {
					continue
				}

				current := dp[i][j]
				visited[i][j] = true
				changed = true

				// Ход 1: 2 столбца вправо, 1 строку вниз
				if i+1 <= N && j+2 <= M {
					dp[i+1][j+2] += current
					visited[i+1][j+2] = false
				}

				// Ход 2: 1 столбец вправо, 2 строки вниз
				if i+2 <= N && j+1 <= M {
					dp[i+2][j+1] += current
					visited[i+2][j+1] = false
				}

				// Ход 3: 2 столбца вправо, 1 строку вверх
				if i-1 >= 1 && j+2 <= M {
					dp[i-1][j+2] += current
					visited[i-1][j+2] = false
				}
			}
		}

		// Если мы достигли конца и изменили что-то, продолжаем
		// Проверяем, нужно ли перезапустить обработку с начала
		if !changed && dp[N][M] > 0 {
			// Проверяем, есть ли необработанные клетки в столбце M
			for i := 1; i <= N; i++ {
				if !visited[i][M] && dp[i][M] > 0 {
					changed = true
					break
				}
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
