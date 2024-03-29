package main

func climbStairsV2(n int) int {
	/*
		METHOD: Recursion
		TIME COMPLEXITY: O(n^2)
		SPACE COMPLEXITY: O(1)
	*/

	// Базовый случай 1: если n равно 1, то есть только один способ подняться - сразу на одну ступеньку.
	if n == 1 {
		return 1
	}
	// Базовый случай 2: если n равно 2, то есть два способа подняться - сначала на одну ступеньку, а затем еще на одну,
	// или сразу на две ступеньки.
	if n == 2 {
		return 2
	}
	// Рекурсивный вызов функции для n-1 ступеней и n-2 ступеней, суммируемый для получения общего количества способов.
	return climbStairsV2(n-1) + climbStairsV2(n-2)
}
