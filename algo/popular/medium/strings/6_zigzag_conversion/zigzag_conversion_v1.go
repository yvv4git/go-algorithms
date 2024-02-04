package __zigzag_conversion

import "strings"

func convertV1(s string, numRows int) string {
	/*
		METHOD: ZigZag
		TIME COMPLEXITY: O(n), внутри основного цикла (for _, c := range s), мы проходим по каждому символу в строке s ровно один раз.
		Space complexity: O(n), внутри алгоритма мы используем слайс строк-строителей (rows) для хранения строк в zigzag-образном порядке.

		Каждый strings.Builder может иметь до n символов (если строка разбивается на одну строку), поэтому общее количество используемой памяти будет не более n.
		Также мы используем дополнительную память для строителя result, который может иметь до n символов.
		Таким образом, общий объем используемой памяти не превышает n, что дает нам сложность O(n).

		Для того чтобы конвертировать строку "PAYPALISHIRING" в zigzag-образ с 3 строками, мы должны следовать следующему алгоритму:
		1. Разделить строку на строки. В этом случае, мы получим 3 строки.
		2. Заполнить каждую строку по одному символу.
		3. Перемещаться по строкам вверх-вниз, добавляя символы.
		4. Объединить все строки вместе.
	*/
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, numRows)
	curRow := 0
	goingDown := false

	for _, c := range s {
		rows[curRow].WriteRune(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}

	return result.String()
}
