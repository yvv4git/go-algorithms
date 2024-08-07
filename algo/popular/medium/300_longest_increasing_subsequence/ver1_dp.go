package main

import "fmt"

// Функция для нахождения длины самой длинной возрастающей подпоследовательности
func lengthOfLIS(nums []int) int {
	/*
		METHOD: Dynamic programming
		В данном решении используется динамическое программирование (Dynamic Programming, DP).
		Основная идея заключается в том, чтобы создать массив dp, где dp[i] представляет длину самой длинной возрастающей подпоследовательности,
		которая заканчивается на элементе nums[i]. Мы инициализируем все элементы dp значением 1,
		так как минимальная длина LIS для любого элемента — это он сам. Затем мы проходим по массиву nums
		и для каждого элемента nums[i] проверяем все предыдущие элементы nums[j] (где j < i). Если nums[i] > nums[j], то обновляем dp[i] как max(dp[i], dp[j] + 1).
		В конце мы возвращаем максимальное значение в массиве dp, которое и будет длиной LIS.

		TIME COMPLEXITY: O(n^2), т.к. внешний цикл проходит по всем элементам массива nums, что дает O(n). Так и внутренний
		цикл также проходит по всем предыдущим элементам для каждого i, что дает O(n). Таким образом, общая временная сложность составляет O(n^2).

		SPACE COMPLEXITY: O(n), т.к. мы используем дополнительный массив dp размером n, что дает пространственную сложность O(n).
	*/
	// Если массив пуст, возвращаем 0
	if len(nums) == 0 {
		return 0
	}

	// Создаем массив dp, где dp[i] представляет длину LIS, заканчивающейся на элементе nums[i]
	dp := make([]int, len(nums))
	// Инициализируем все элементы dp значением 1, так как минимальная длина LIS для любого элемента - это он сам
	for i := range dp {
		dp[i] = 1
	}

	// Максимальная длина LIS изначально равна 1
	maxLIS := 1

	// Проходим по массиву nums
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			// Если nums[i] > nums[j], это означает, что мы можем продолжить LIS, заканчивающуюся на nums[j]
			if nums[i] > nums[j] {
				// Обновляем dp[i] как максимум из текущего dp[i] и dp[j] + 1
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// Обновляем максимальную длину LIS
		maxLIS = max(maxLIS, dp[i])
	}

	return maxLIS
}

// Вспомогательная функция для нахождения максимума двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Length of LIS:", lengthOfLIS(nums)) // Вывод: Length of LIS: 4
}
