package _518_Coin_Change2

func changeV2(amount int, coins []int) int {
	/*
		TASK: Подсчитать количество комбинаций, которыми можно составить сумму amount,
		используя монеты из coins неограниченное количество раз.
		Порядок монет не имеет значения (комбинации, не перестановки).
		Если сумму нельзя составить — вернуть 0.

		APPROACH: Recursion with Memoization (Top-down DP)
		1. Рекурсивно перебираем монеты: на каждом шаге либо берём текущую монету,
		   либо пропускаем её (переходим к следующей).
		2. Если remaining == 0 — нашли комбинацию, возвращаем 1.
		3. Если remaining < 0 или монеты кончились — возвращаем 0.
		4. Кешируем результат для пары (i, remaining), чтобы избежать повторных вычислений.

		TIME COMPLEXITY: O(N × M)
		- N = количество монет, M = amount.
		- Каждая пара (i, remaining) вычисляется один раз.

		SPACE COMPLEXITY: O(N × M)
		- Храним memo-карту с N × M записями в худшем случае.
		- Глубина рекурсии до N.
	*/
	memo := make(map[[2]int]int)

	var dfs func(i, remaining int) int
	dfs = func(i, remaining int) int {
		if remaining == 0 {
			return 1
		}
		if i == len(coins) || remaining < 0 {
			return 0
		}

		key := [2]int{i, remaining}
		if v, ok := memo[key]; ok {
			return v
		}

		skip := dfs(i+1, remaining)
		take := dfs(i, remaining-coins[i])

		memo[key] = skip + take
		return memo[key]
	}

	return dfs(0, amount)
}
