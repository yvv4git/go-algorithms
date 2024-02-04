package _335_minimum_difficulty_of_a_job_schedule

func minDifficultyV1(jobDifficulty []int, d int) int {
	/*
		METHOD: Dynamic programming
		TIME COMPLEXITY: O(n^2 * d), где n - количество задач, d - количество дней.
		SPACE COMPLEXITY: O(n * d), так как мы используем дополнительный массив dp для хранения результатов
	*/

	// Длина массива сложностей задач.
	length := len(jobDifficulty)

	// Если количество задач меньше, чем количество дней, то невозможно выполнить все задачи.
	if length < d {
		return -1
	}

	// Создаем двумерный массив dp для хранения результатов.
	dp := make([][]int, d)
	for i := range dp {
		dp[i] = make([]int, length)
	}

	// Базовый случай: минимальная сложность для первого дня равна сложности первой задачи.
	dp[0][0] = jobDifficulty[0]
	for i := 1; i < length; i++ {
		// Для каждого дня, минимальная сложность равна максимуму между предыдущей минимальной сложностью и сложностью текущей задачи.
		dp[0][i] = max(dp[0][i-1], jobDifficulty[i])
	}

	// Для каждого дня, минимальная сложность равна минимальному значению между предыдущей минимальной сложностью и сложностью текущей задачи.
	for i := 1; i < d; i++ {
		for j := i; j < length; j++ {
			dp[i][j] = dp[i-1][j-1] + jobDifficulty[j]
			maxDiff := jobDifficulty[j]
			for k := j - 1; k >= i-1; k-- {
				maxDiff = max(maxDiff, jobDifficulty[k+1])
				dp[i][j] = min(dp[i][j], dp[i-1][k]+maxDiff)
			}
		}
	}

	// Возвращаем минимальную сложность для последнего дня.
	return dp[d-1][length-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
