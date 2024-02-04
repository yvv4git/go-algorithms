package main

import (
	"math"
	"strings"
)

func myAtoiV4(s string) int {
	/*
		METHOD: Iteration
		Time complexity: O(n)
		Space complexity: O(1) - в общем случае, в худшем случае, если строка состоит только из цифр, то n записей придется копировать.
	*/

	// Инициализируем переменную result, в которой будет храниться результат преобразования.
	var result int
	// Инициализируем переменную modification, которая определяет знак числа.
	modification := 1

	// Проходим по строке s, используя функцию Trim для удаления пробелов в начале и конце строки.
	for index, data := range strings.Trim(s, " ") {
		// Если символ является цифрой, то преобразуем его в число и добавляем к результату.
		if data >= '0' && data <= '9' {
			result = result*10 + int(data-'0')
			// Проверяем, не выходит ли результат за пределы диапазона int32.
			switch {
			case result*modification >= math.MaxInt32:
				return math.MaxInt32
			case result*modification <= math.MinInt32:
				return math.MinInt32
			}
		} else {
			// Если символ не является цифрой, то проверяем, является ли он знаком минус или плюс.
			if index == 0 {
				switch data {
				case '-':
					modification = -1
					continue
				case '+':
					modification = 1
					continue
				default:
					return 0
				}
			} else {
				// Если символ не является цифрой и не является знаком минус или плюс, то возвращаем результат.
				return result * modification
			}
		}
	}

	// Возвращаем результат преобразования.
	return result * modification
}
