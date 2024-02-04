package sum_of_digits_float

import "math"

func sumDigits(n float64) int {
	/*
		METHOD: Math + loop
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)
	*/
	sum := 0
	whole, fraction := math.Modf(n)
	// whole - целая часть
	// fraction - дробная часть.

	// Пока целая часть больше нуля, вычисляем сумму чисел в ней.
	for whole > 0 {
		sum += int(whole) % 10
		whole /= 10
	}

	// Отдельно вычисляем сумму чисел в дробной части.
	for fraction > 0 {
		fraction *= 10
		sum += int(fraction)
		fraction -= float64(int(fraction))
	}

	return sum
}
