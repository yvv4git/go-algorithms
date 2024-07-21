package main

func coinChangeV2(coins []int, amount int) int {
	/*
		METHOD: Dynamic programming with recursion

		TIME COMPLEXITY: O(n * k), где n — это сумма `amount`, а k — количество различных номиналов монет.

		SPACE COMPLEXITY: O(n), где n — это сумма `amount`. Мы используем массив `dp` размером `amount + 1`.
	*/
	var c func(coins []int, amount int, dp map[int]int) int
	var r func(coins []int, amount int) int

	c = func(coins []int, amount int, dp map[int]int) int {
		// Базовый случай: если сумма равна 0, то не нужно ни одной монеты.
		if amount == 0 {
			return 0
		}
		// Базовый случай: если сумма меньше 0, то возвращаем -1, так как это недопустимое состояние.
		if amount < 0 {
			return -1
		}
		// Проверяем, есть ли уже вычисленное значение для данной суммы в мапе `dp`.
		v, ok := dp[amount]
		if ok {
			return v
		}
		// Инициализируем переменную для хранения минимального количества монет.
		m := -1
		// Перебираем все номиналы монет.
		for i := 0; i < len(coins); i++ {
			// Рекурсивно вычисляем минимальное количество монет для суммы `amount - coins[i]`.
			tmp := c(coins, amount-coins[i], dp)
			// Если результат не -1, то обновляем минимальное количество монет.
			if tmp != -1 {
				if m == -1 {
					m = 1 + tmp
				} else {
					m = min(m, 1+tmp)
				}
			}
		}
		// Сохраняем вычисленное значение в мапе `dp`.
		dp[amount] = m
		return m
	}
	r = func(coins []int, amount int) int {
		// Создаем массив `dp` размером `amount + 1` и заполняем его максимальными значениями.
		dp := make([]int, amount+1)
		for i := 1; i <= amount; i++ {
			dp[i] = 1 << 31
			// Перебираем все номиналы монет.
			for j := 0; j < len(coins); j++ {
				// Если номинал монеты меньше или равен текущей сумме и предыдущее значение меньше текущего,
				// то обновляем текущее значение.
				if coins[j] <= i && dp[i-coins[j]] < dp[i] {
					dp[i] = 1 + dp[i-coins[j]]
				}
			}
		}
		// Если значение для суммы `amount` осталось максимальным, то возвращаем -1.
		if dp[amount] == 1<<31 {
			return -1
		}
		return dp[amount]
	}

	//return c(coins, amount, map[int]int{})
	return r(coins, amount)
}

//func min(a, b int) int {
//	if a == -1 {
//		return b
//	}
//	if b == -1 {
//		return a
//	}
//	if a < b {
//		return a
//	}
//	return b
//}
