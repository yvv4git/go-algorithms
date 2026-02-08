package main

import (
	"bufio"
	"fmt"
	"os"
)

func rob(nums []int) int {
	/*
		INTUITION:
		Дома расположены по кругу, поэтому первый и последний дом - соседи.
		Нельзя грабить оба конца одновременно, иначе сработает сигнализация.

		Решение: разбиваем задачу на две независимые:
		1. Грабим дома [0...n-2] (исключаем последний)
		2. Грабим дома [1...n-1] (исключаем первый)
		Максимум из этих двух вариантов - ответ.

		Для каждого варианта используем классический House Robber:
		- dp[i] = максимальная сумма для домов [0...i]
		- Переход: dp[i] = max(dp[i-1], dp[i-2] + nums[i])
		- Либо пропускаем текущий дом, либо грабим его (тогда пропускаем предыдущий)

		METHOD: Dynamic Programming (разделение на два случая)
		TIME COMPLEXITY: O(n) - один проход для каждого случая
		SPACE COMPLEXITY: O(1) - используем только две переменные
	*/

	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	// robRange возвращает максимальную сумму для диапазона [start...end]
	robRange := func(start, end int) int {
		prev1 := 0 // dp[i-1]
		prev2 := 0 // dp[i-2]

		for i := start; i <= end; i++ {
			current := max(prev1, prev2+nums[i])
			prev2 = prev1
			prev1 = current
		}
		return prev1
	}

	// Случай 1: грабим дома с 0 до n-2 (исключаем последний)
	// Случай 2: грабим дома с 1 до n-1 (исключаем первый)
	return max(robRange(0, n-2), robRange(1, n-1))
}

// max возвращает максимум из двух целых чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)

	// Читаем количество домов
	var n int
	fmt.Fscan(in, &n)

	// Читаем суммы денег в каждом доме
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &nums[i])
	}

	result := rob(nums)
	fmt.Println(result)
}
