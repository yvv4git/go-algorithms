package main

import "strconv"

func countNumbersWithUniqueDigitsV2(n int) int {
	/*
		METHOD: Memoization with recursion
		Для решения задачи с использованием рекурсии с мемоизацией,
		мы можем использовать массив для хранения уже вычисленных значений, чтобы избежать повторных вычислений.

		TIME COMPLEXITY: O(n), так как мы проходимся по всем числам от 0 до n.

		SPACE COMPLEXITY: O(n), так как мы используем дополнительный массив для хранения уже вычисленных значений.
	*/
	if n == 0 {
		return 1
	}

	// Инициализируем массив dp, где dp[i] будет содержать количество уникальных чисел, которые могут быть получены из i цифр
	dp := make([]int, n+1)

	// Инициализируем массив used, где used[i] будет содержать информацию о том, использовалась ли цифра i
	used := make([]bool, 10)

	// Вызов рекурсивной функции для вычисления количества уникальных чисел
	return countNumbersWithUniqueDigitsRecursive(n, 0, true, true, dp, used)
}

func countNumbersWithUniqueDigitsRecursive(n int, pos int, limited bool, lead bool, dp []int, used []bool) int {
	// Если мы достигли конца числа, возвращаем 1, так как это означает, что мы нашли уникальное число
	if pos == n {
		return 1
	}

	// Если мы уже вычислили это значение, возвращаем его
	if !limited && lead && dp[pos] != 0 {
		return dp[pos]
	}

	var res int
	var limit int

	if limited {
		limit = int([]byte(strconv.Itoa(n))[pos] - '0')
	} else {
		limit = 9
	}

	for i := 0; i <= limit; i++ {
		if lead && i == 0 {
			res += countNumbersWithUniqueDigitsRecursive(n, pos+1, limited && i == limit, true, dp, used)
		} else if !used[i] {
			used[i] = true
			res += countNumbersWithUniqueDigitsRecursive(n, pos+1, limited && i == limit, false, dp, used)
			used[i] = false
		}
	}

	// Если мы не достигли ограничения, сохраняем результат в dp
	if !limited && lead {
		dp[pos] = res
	}

	return res
}
