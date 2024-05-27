package main

import (
	"fmt"
	"math"
)

// Функция для поиска максимального произведения трех чисел
func maximumProduct(nums []int) int {
	/*
		METHOD: Iterate
		Мы используем подход, основанный на поиске трех наибольших и двух наименьших чисел,
		что позволяет нам найти максимальное произведение трех чисел в массиве за один проход.

		TIME COMPLEXITY: O(n), где n - количество чисел в массиве, потому что мы проходим по массиву всего один раз.

		SPACE COMPLEXITY: O(1), так как мы используем фиксированное количество переменных, не зависящих от размера входных данных.
	*/
	// Инициализируем переменные для трех наибольших чисел и двух наименьших
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	// Проходим по всем числам в массиве
	for _, num := range nums {
		// Обновляем наименьшие числа
		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num <= min2 {
			min2 = num
		}

		// Обновляем наибольшие числа
		if num >= max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max3 = max2
			max2 = num
		} else if num >= max3 {
			max3 = num
		}
	}

	// Возвращаем максимум из произведений двух наибольших и одного наименьшего
	// числа и трех наибольших чисел
	return max(min1*min2*max1, max1*max2*max3)
}

// Функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример использования функции
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumProduct(nums)) // Вывод: 60
}
