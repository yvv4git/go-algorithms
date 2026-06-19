package _518_Coin_Change2

func changeV3(amount int, coins []int) int {
	/*
		TASK: Подсчитать количество комбинаций, которыми можно составить сумму amount,
		используя монеты из coins неограниченное количество раз.
		Порядок монет не имеет значения (комбинации, не перестановки).
		Если сумму нельзя составить — вернуть 0.

		APPROACH: 2D Dynamic Programming
		Используем двумерный массив dp[i][a], где:
		- i — количество рассмотренных монет (0..len(coins))
		- a — целевая сумма (0..amount)
		dp[i][a] = количество комбинаций собрать сумму a,
		используя первые i типов монет.

		База: dp[0][0] = 1 (пустая комбинация)
		Для каждого i от 1 до len(coins):
		  Для каждой суммы a от 0 до amount:
		    dp[i][a] = dp[i-1][a]                  // пропустить монету
		             + dp[i][a-coins[i-1]]         // взять монету (если a >= coins[i-1])

		Это явно показывает два выбора: "skip or take".

		TIME COMPLEXITY: O(N × M)
		- N = len(coins), M = amount.
		- Два вложенных цикла.

		SPACE COMPLEXITY: O(N × M)
		- Храним таблицу размером (N+1) × (M+1).
		- Можно оптимизировать до O(M), как в changeV1.
	*/

	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}

	dp[0][0] = 1

	for i := 1; i <= n; i++ {
		for a := 0; a <= amount; a++ {
			dp[i][a] = dp[i-1][a]
			if a >= coins[i-1] {
				dp[i][a] += dp[i][a-coins[i-1]]
			}
		}
	}

	return dp[n][amount]
}
