package main

func sumOfCompatibleNumbersV3(n, k int) int {
	/*
		TASK: Найти сумму всех положительных целых x, удовлетворяющих:
		abs(n - x) <= k и (n & x) == 0.

		METHOD: Brute Force
		Перебираем все x от max(1, n-k) до n+k и проверяем оба условия.

		TIME COMPLEXITY: O(k)
		- k <= 100 (constraints), максимум 201 итерация.

		SPACE COMPLEXITY: O(1)
		- Используем только одну переменную для суммы.
	*/
	L := max(1, n-k)
	R := n + k

	totalSum := 0

	for x := L; x <= R; x++ {
		if n&x == 0 {
			totalSum += x
		}
	}

	return totalSum
}
