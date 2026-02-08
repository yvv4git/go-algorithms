//go:build ignore

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

		Для каждого варианта используем классический House Robber с одномерным DP:
		- dp[i] = максимальная сумма для домов [start...i]
		- Переход: dp[i] = max(dp[i-1], dp[i-2] + nums[i])
		- Либо пропускаем текущий дом, либо грабим его (тогда пропускаем предыдущий)

		METHOD: Dynamic Programming (одномерный массив)
		TIME COMPLEXITY: O(n) - один проход для каждого случая
		SPACE COMPLEXITY: O(n) - хранение dp-таблицы
	*/

	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// robRange возвращает максимальную сумму для диапазона [start...end]
	// с использованием одномерного массива dp
	robRange := func(start, end int) int {
		length := end - start + 1
		if length <= 0 {
			return 0
		}

		// dp[i] - максимальная сумма для домов [start...start+i]
		dp := make([]int, length)

		// Базовые случаи
		dp[0] = nums[start]

		if length > 1 {
			dp[1] = max(nums[start], nums[start+1])
		}

		// Заполняем dp массив
		for i := 2; i < length; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[start+i])
		}

		return dp[length-1]
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
