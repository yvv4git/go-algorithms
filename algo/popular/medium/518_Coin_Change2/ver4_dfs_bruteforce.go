package _518_Coin_Change2

func changeV4(amount int, coins []int) int {
	/*
		TASK: Подсчитать количество комбинаций, которыми можно составить сумму amount,
		используя монеты из coins неограниченное количество раз.
		Порядок монет не имеет значения (комбинации, не перестановки).
		Если сумму нельзя составить — вернуть 0.

		APPROACH: DFS Brute Force (без мемоизации)
		Рекурсивно перебираем все варианты: на каждом шаге либо берём текущую монету,
		либо пропускаем. Этот подход экспоненциальный — нужен только для демонстрации,
		почему требуется DP.

		TIME COMPLEXITY: O(2^N) — экспоненциальный, каждую монету можно
		брать или не брать многократно.

		SPACE COMPLEXITY: O(N) — глубина рекурсии.
	*/

	var dfs func(i, remaining int) int
	dfs = func(i, remaining int) int {
		if remaining == 0 {
			return 1
		}
		if i == len(coins) || remaining < 0 {
			return 0
		}

		skip := dfs(i+1, remaining)
		take := dfs(i, remaining-coins[i])

		return skip + take
	}

	return dfs(0, amount)
}
