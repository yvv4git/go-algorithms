package __zigzag_conversion

func convertV2(s string, numRows int) string {
	/*
		METHOD: Matrix
		TIME COMPLEXITY: O(n)
		Space complexity: O(n)

		Time Complexity для этого алгоритма составляет O(n), где n - длина входной строки.
		Внутри алгоритма мы проходимся по строке и по матрице. Проход по строке занимает O(n) времени,
		а проход по матрице также занимает O(n) времени, так как мы проходимся по каждой ячейке.
		Таким образом, общий временный сложность составляет O(n).

		Space Complexity для этого алгоритма составляет O(n), где n - длина входной строки.
		Внутри алгоритма мы используем матрицу для хранения результата. Матрица имеет размер numRows x len(s),
		где len(s) - длина входной строки. Таким образом, общий объем используемой памяти не превышает n,
		что дает нам сложность O(n).
	*/
	if numRows == 1 {
		return s
	}

	matrix := make([][]rune, numRows)
	for i := range matrix {
		matrix[i] = make([]rune, len(s))
	}

	row, col := 0, 0
	goingDown := false

	for _, c := range s {
		matrix[row][col] = c
		if row == numRows-1 || row == 0 {
			goingDown = !goingDown
		}
		if goingDown {
			row++
		} else {
			row--
			col++
		}
	}

	result := make([]rune, 0, len(s))
	for _, row := range matrix {
		for _, c := range row {
			if c != 0 {
				result = append(result, c)
			}
		}
	}

	return string(result)
}
