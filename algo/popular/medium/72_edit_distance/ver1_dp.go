package main

import (
	"fmt"
	"math"
)

// Функция minDistance вычисляет минимальное количество операций редактирования,
// необходимых для превращения одной строки в другую.
func minDistance(word1 string, word2 string) int {
	/*
		METHOD: Dynamic programming
		Данное решение использует динамическое программирование, что означает,
		что мы разбиваем исходную задачу на более мелкие подзадачи и решаем их по очереди, сохраняя результаты для последующих вычислений.
		В данном случае, мы создаем матрицу dp, где dp[i][j] представляет минимальное количество операций,
		необходимых для превращения префикса первой строки длиной i в префикс второй строки длиной j.

		TIME COMPLEXITY: O(m*n), где m и n - длины двух строк. Это обусловлено тем, что мы проходим по каждому символу каждой строки ровно один раз.

		SPACE COMPLEXITY: O(m*n), так как мы используем дополнительную матрицу dp размером mn для хранения промежуточных результатов.
	*/
	// Длина первой и второй строк.
	m, n := len(word1), len(word2)

	// Создаем двумерный массив dp размером (m+1)x(n+1), где dp[i][j] будет
	// представлять минимальное количество операций, необходимых для превращения
	// префикса первой строки длиной i в префикс второй строки длиной j.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Инициализируем первую строку и первый столбец массива dp.
	// dp[i][0] будет равно i, так как для превращения префикса первой строки длиной i
	// в пустую строку нужно i операций (все удалить). Аналогично для dp[0][j].
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	// Заполняем оставшуюся часть массива dp.
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// Если символы в обеих строках одинаковые, то количество операций не изменится.
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// Если символы разные, то находим минимальное значение из трех возможных вариантов:
				// 1) Заменить символ в первой строке на символ из второй строки (dp[i-1][j-1] + 1).
				// 2) Удалить символ из первой строки (dp[i-1][j] + 1).
				// 3) Вставить символ из второй строки в первую (dp[i][j-1] + 1).
				dp[i][j] = int(math.Min(float64(dp[i-1][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i][j-1]))) + 1)
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(minDistance("kitten", "sitting"))
}
