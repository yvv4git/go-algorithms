package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
	/*
		METHOD: Dynamic Programming
		1. Группируем числа по значению и считаем сумму для каждого уникального числа.
		2. Сортируем уникальные числа.
		3. Применяем DP: dp[i] - максимальная сумма для первых i уникальных чисел,
		   учитывая, что нельзя брать числа, отличающиеся на 1 (как в House Robber).

		TIME COMPLEXITY: O(n log n) (сортировка уникальных чисел, n - длина массива)
		SPACE COMPLEXITY: O(n) (map для сумм и dp массив)
	*/
	if len(nums) == 0 {
		return 0
	}

	// Группируем числа по значению и считаем сумму
	sumMap := make(map[int]int)
	for _, num := range nums {
		sumMap[num] += num
	}

	// Получаем отсортированный список уникальных чисел
	unique := make([]int, 0, len(sumMap))
	for num := range sumMap {
		unique = append(unique, num)
	}
	sort.Ints(unique)

	// DP: dp[i] - максимум для первых i уникальных чисел
	n := len(unique)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return sumMap[unique[0]]
	}

	dp := make([]int, n)
	dp[0] = sumMap[unique[0]]
	if unique[1] == unique[0]+1 {
		dp[1] = max(dp[0], sumMap[unique[1]])
	} else {
		dp[1] = dp[0] + sumMap[unique[1]]
	}

	for i := 2; i < n; i++ {
		if unique[i] == unique[i-1]+1 {
			dp[i] = max(dp[i-1], dp[i-2]+sumMap[unique[i]])
		} else {
			dp[i] = dp[i-1] + sumMap[unique[i]]
		}
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Примеры
	fmt.Println(deleteAndEarn([]int{3, 4, 2}))          // 6
	fmt.Println(deleteAndEarn([]int{2, 2, 3, 3, 3, 4})) // 9
}
