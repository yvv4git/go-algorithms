package main

import (
	"fmt"
	"sort"
)

// Функция для вычисления Манхэттенского расстояния
func manhattanDistance(r1, c1, r2, c2 int) int {
	return abs(r1-r2) + abs(c1-c2)
}

// Функция для вычисления абсолютного значения
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Функция для сортировки ячеек матрицы по расстоянию до центральной ячейки
func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
	/*
		METHOD: Sort by key
		Используемый подход в этом решении называется "сортировка по ключу".
		Мы сортируем массив ячеек матрицы на основе некоторого ключа, в данном случае - расстояния до центральной ячейки.
		Этот подход широко используется в различных задачах, где требуется упорядочить элементы по некоторому критерию.

		TIME COMPLEXITY: O(RClog(RC)), где rows - количество строк, cols - количество столбцов.
		Это связано с тем, что мы генерируем и сортируем все ячейки матрицы.

		SPACE COMPLEXITY: O(RC), так как мы храним все ячейки в массиве.
		Это связано с тем, что для хранения всех ячеек матрицы нам нужен массив размера rows*cols.
	*/
	// Создаем массив для хранения всех ячеек
	cells := make([][]int, 0, rows*cols)

	// Заполняем массив ячеек
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cells = append(cells, []int{r, c})
		}
	}

	// Сортируем ячейки по расстоянию до центральной ячейки
	sort.Slice(cells, func(i, j int) bool {
		d1 := manhattanDistance(cells[i][0], cells[i][1], rCenter, cCenter)
		d2 := manhattanDistance(cells[j][0], cells[j][1], rCenter, cCenter)
		// Если расстояния равны, сортируем по координатам
		if d1 == d2 {
			return cells[i][0] < cells[j][0] || cells[i][0] == cells[j][0] && cells[i][1] < cells[j][1]
		}
		return d1 < d2
	})

	return cells
}

func main() {
	R := 2
	C := 3
	r0 := 1
	c0 := 2

	cells := allCellsDistOrder(R, C, r0, c0)
	for _, cell := range cells {
		fmt.Println(cell)
	}
}
