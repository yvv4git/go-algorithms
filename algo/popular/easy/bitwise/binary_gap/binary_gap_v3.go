package binary_gap

import "math"

func binaryGapV3(n int) int {
	/*
		METHOD: Bitwise
		Time complexity: O(n)
		Space complexity: O(1)
	*/
	interval, result := 0, math.MinInt32

	for n > 0 {
		if n&1 == 0 { // Если младший(правый) разряд n содержит 1, то делаем right shift. Еще это эквивалентно делению на 2
			n = n >> 1
		} else { // Если младший(правый) разряд 0, пытаемся посчитать дистанцию
			interval = 0
			n = n >> 1 // Снова двигаем вправо(right shift)
			for n > 0 {
				interval++
				if n&1 == 1 { // Проверяем младший(правый) разряд на 1
					break
				}
				n = n >> 1
			}

			if result < interval {
				result = interval
			}
		}
	}

	return result
}
