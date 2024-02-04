package main

import (
	"fmt"
	"math"
)

func titleToNumber(columnTitle string) int {
	/*
		METHOD:
		TIME COMPLEXITY: O(n), где n - количество символов в названии столбца. Это обусловлено тем, что мы проходим по каждому символу в строке ровно один раз.
		Space complexity: O(1), так как мы используем фиксированное количество переменных, не зависящих от размера входных данных.

		Для решения этой задачи мы можем использовать идею, что каждый символ столбца Excel представляет собой позицию в алфавите,
		где A - 1, B - 2, ..., Z - 26. Затем мы можем использовать эту информацию для вычисления порядкового номера столбца,
		как если бы это была запись в системе счисления с основанием 26.
	*/
	columnNumber := 0
	length := len(columnTitle)

	for i := 0; i < length; i++ {
		// Вычисляем значение текущего символа в системе счисления с основанием 26
		charValue := int(columnTitle[i] - 'A' + 1)
		// Добавляем значение символа, умноженное на его позицию (с учетом нулевой позиции)
		columnNumber += charValue * int(math.Pow(26, float64(length-i-1)))
	}

	return columnNumber

}

func main() {
	fmt.Println(titleToNumber("ZY")) // Выведет: 701
}
