package __zigzag_conversion

import "strings"

func convertV3(s string, numRows int) string {
	/*
		Method: Simple loop
		Time complexity: O(n)
		Space complexity: O(n)

		Команда curRow-- используется для перемещения вверх по строкам.
		В "zigzag" формате, когда мы достигаем нижней границы строк, мы начинаем двигаться вверх.
		Команда curRow-- уменьшает текущую строку на 1, что позволяет нам перемещаться вверх по строкам.

		Например, если у нас есть 3 строки, тогда curRow будет иметь значения 0, 1, 2.
		Когда curRow станет равным 2 (т.е., мы достигли нижней границы), мы начнем двигаться вверх.
		Для этого мы уменьшаем curRow на 1, чтобы переместиться на строку выше.

		Таким образом, curRow-- используется для перемещения вверх по строкам в "zigzag" формате.
	*/
	if numRows == 1 { // Можно вернуть результат, так как надо разбить текст на строки
		return s
	}

	rows := make([]string, numRows) // Нужен массив строк размером numRows
	curRow := 0
	goingDown := false

	for _, c := range s { // Проходим по каждому символу в строке
		rows[curRow] += string(c)               // Добавляем символ к результирующей строке
		if curRow == 0 || curRow == numRows-1 { //  Проверяем, достигнута ли нижняя или верхняя граница массива строк
			goingDown = !goingDown // Если достигнута - выставляем флаг в true
		}

		if goingDown {
			curRow++ // Переходим на следующую строку
		} else {
			curRow--
		}
	}

	return strings.Join(rows, "")
}
