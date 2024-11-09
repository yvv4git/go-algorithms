package main

import (
	"fmt"
	"math"
)

func smallestRangeI(nums []int, k int) int {
	/*
		METHOD: Linear Search with Constant Space
		Этот метод основан на математическом подходе для нахождения минимально возможного диапазона
		между максимальным и минимальным элементами массива после применения операции, которая
		позволяет увеличить или уменьшить каждый элемент на любое значение в диапазоне от -k до k.

		TIME COMPLEXITY: O(n)
		O(n) - где n — количество элементов в массиве. Мы проходим по массиву один раз, чтобы найти
		максимальное и минимальное значения.

		SPACE COMPLEXITY: O(1)
		O(1) - мы используем фиксированное количество дополнительной памяти (переменные maxNum, minNum,
		newMax, newMin и rangeDiff), которое не зависит от размера входного массива.
	*/
	// Находим максимальное и минимальное значения в массиве
	maxNum := math.MinInt32
	minNum := math.MaxInt32

	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}

		if num < minNum {
			minNum = num
		}
	}

	// Вычисляем новые максимальное и минимальное значения после применения операции
	newMax := maxNum - k
	newMin := minNum + k

	// Вычисляем минимальный диапазон
	rangeDiff := newMax - newMin

	// Если диапазон отрицательный, возвращаем 0
	if rangeDiff < 0 {
		return 0
	}

	// Возвращаем минимальный диапазон
	return rangeDiff
}

func main() {
	// Пример использования
	nums := []int{1, 3, 6}
	k := 3
	result := smallestRangeI(nums, k)
	fmt.Println(result) // Ожидаемый вывод: 0
}
