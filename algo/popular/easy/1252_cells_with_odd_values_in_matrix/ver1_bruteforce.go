package main

import "fmt"

func oddCells(m int, n int, indices [][]int) int {
	/*
		APPROACH: Brute-froce
		Используем два массива: один для подсчета изменений в строках, другой — в столбцах.
		Для каждого индекса увеличиваем соответствующие элементы в массивах rows и cols.
		После этого считаем количество нечетных значений в строках и столбцах.
		Итоговое количество нечетных ячеек вычисляется по формуле:
		oddRows * n + oddCols * m - 2 * oddRows * oddCols.

		TIME COMPLEXITY: O(m + n + len(indices)),
		где m — количество строк, n — количество столбцов,
		len(indices) — количество операций инкремента.

		SPACE COMPLEXITY: O(m + n),
		так как используются два массива для хранения изменений в строках и столбцах.
	*/
	rows := make([]int, m)
	cols := make([]int, n)

	for _, index := range indices {
		rows[index[0]]++
		cols[index[1]]++
	}

	oddRows := 0
	for _, r := range rows {
		if r%2 == 1 {
			oddRows++
		}
	}

	oddCols := 0
	for _, c := range cols {
		if c%2 == 1 {
			oddCols++
		}
	}

	return oddRows*n + oddCols*m - 2*oddRows*oddCols
}

func main() {
	m, n := 2, 3
	indices := [][]int{{0, 1}, {1, 1}}
	fmt.Println(oddCells(m, n, indices)) // Ожидаемый вывод: 6
}
