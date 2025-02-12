package main

import (
	"fmt"
	"math"
)

// Функция minStartValue находит минимальное начальное значение startValue
// так, чтобы пошаговая сумма оставалась положительной.
func minStartValue(nums []int) int {
	/*
		METHOD: Linear Scan (Линейный проход по массиву)
		- Мы последовательно вычисляем префиксную сумму (step-by-step sum).
		- Следим за минимальной промежуточной суммой.
		- Итоговое значение startValue вычисляется по формуле: 1 - minPrefixSum.

		TIME COMPLEXITY: O(n)
		- Один проход по массиву (линейное время).

		SPACE COMPLEXITY: O(1)
		- Используются только несколько переменных (константная память).
	*/
	minSum := math.MaxInt32 // Переменная для хранения минимальной префиксной суммы
	currentSum := 0         // Текущая сумма (префиксная сумма)

	// Проходим по массиву, вычисляя префиксную сумму и минимальное значение
	for _, num := range nums {
		currentSum += num
		if currentSum < minSum {
			minSum = currentSum
		}
	}

	// Минимально допустимое startValue — 1 - minSum (должно быть ≥ 1)
	return max(1-minSum, 1)
}

// Вспомогательная функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример использования функции
	nums := []int{-3, 2, -3, 4, 2}
	fmt.Println(minStartValue(nums)) // Ожидаемый результат: 5
}
