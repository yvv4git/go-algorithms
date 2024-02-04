package _264_largest_3_same_digit_number_in_string

import (
	"fmt"
	"strings"
)

// largestGoodInteger находит наибольшее число, состоящее из одной и той же цифры, длиной 3 символа.
// Функция принимает строку num, представляющую большое целое число.
func largestGoodIntegerV1(num string) string {
	/*
		METHOD: Enumeration / Перебор
		TIME COMPLEXITY: O(n)
		Space complexity: O(1)

		TIME COMPLEXITY:
		Временная сложность этого алгоритма составляет O(n), где n - длина входной строки,
		потому что мы проходим по каждому символу в строке.

		Space complexity:
		Пространственная сложность составляет O(1), поскольку мы используем фиксированное количество дополнительной памяти,
		независимо от размера входной строки.
	*/

	// Проходим цифры от 9 до 0.
	for i := 9; i >= 0; i-- {
		// Формируем строку, состоящую из одной и той же цифры, повторяющейся 3 раза.
		digit := strings.Repeat(fmt.Sprintf("%d", i), 3)
		// Проверяем, является ли строка digit подстрокой строки num.
		if strings.Contains(num, digit) {
			// Если да, возвращаем цифру.
			return digit
		}
	}

	// Если такой цифры не найдено, возвращаем пустую строку.
	return ""
}
