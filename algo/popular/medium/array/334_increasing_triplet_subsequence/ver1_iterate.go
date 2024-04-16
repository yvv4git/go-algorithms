package main

import "math"

// Функция increasingTriplet проверяет, существует ли в массиве nums подпоследовательность из трех элементов, которая строго возрастает.
func increasingTriplet(nums []int) bool {
	/*
		METHOD: Iterate
		Используемый подход для решения задачи - это один проход по массиву с целью отслеживания минимальных элементов,
		которые могут стать началом подпоследовательности. Если на каком-то этапе мы находим элемент,
		который больше второго элемента подпоследовательности, то мы нашли подходящую подпоследовательность.

		TIME COMPLEXITY: O(n), где n - количество элементов в массиве, так как мы проходим по массиву всего один раз.

		SPACE COMPLEXITY: O(1), так как мы используем только две дополнительные переменные для отслеживания минимальных элементов подпоследовательности, независимо от размера входного массива.
	*/
	// Если длина массива меньше 3, то подпоследовательность невозможна.
	if len(nums) < 3 {
		return false
	}

	// Инициализация двух переменных first и second на максимальное значение int32.
	// Они будут хранить минимальные элементы, которые могут быть началом подпоследовательности.
	first, second := math.MaxInt32, math.MaxInt32

	// Проход по всем элементам массива nums.
	for _, num := range nums {
		// Если текущий элемент меньше или равен first, обновляем first.
		if num <= first {
			first = num
		} else if num <= second {
			// Если текущий элемент больше first, но меньше или равен second, обновляем second.
			second = num
		} else {
			// Если текущий элемент больше second, значит, мы нашли подпоследовательность.
			return true
		}
	}

	// Если мы прошли весь массив и не нашли подходящую подпоследовательность, возвращаем false.
	return false
}
