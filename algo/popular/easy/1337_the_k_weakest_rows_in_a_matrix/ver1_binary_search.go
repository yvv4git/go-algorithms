package main

import (
	"fmt"
	"sort"
)

// Функция для поиска k самых слабых строк в матрице
func kWeakestRows(mat [][]int, k int) []int {
	/*
		METHOD: Binary Search + Sorting
		1. Создаем слайс для хранения количества единиц и индексов строк.
		2. Проходим по каждой строке матрицы.
		3. Используем бинарный поиск для подсчета единиц (так как строки отсортированы).
		4. Сохраняем количество единиц и индекс строки.
		5. Сортируем строки по количеству единиц, а затем по индексу.
		6. Возвращаем индексы k самых слабых строк.

		TIME COMPLEXITY: O(m * log n + m log m)
		- m * log n: бинарный поиск для каждой строки (m строк, каждая длиной n).
		- m log m: сортировка m строк по количеству единиц и индексам.

		SPACE COMPLEXITY: O(m)
		- Для хранения информации о количестве единиц и индексах строк.
	*/
	// Создаем слайс для хранения количества единиц и индексов строк
	type rowStrength struct {
		strength int // Количество единиц в строке
		index    int // Индекс строки
	}
	strengths := make([]rowStrength, len(mat))

	// Проходим по каждой строке матрицы
	for i, row := range mat {
		// Используем бинарный поиск для подсчета единиц
		left, right := 0, len(row)
		for left < right {
			mid := (left + right) / 2
			if row[mid] == 1 {
				left = mid + 1
			} else {
				right = mid
			}
		}
		// Сохраняем количество единиц и индекс строки
		strengths[i] = rowStrength{strength: left, index: i}
	}

	// Сортируем строки по количеству единиц, а затем по индексу
	sort.Slice(strengths, func(i, j int) bool {
		if strengths[i].strength == strengths[j].strength {
			return strengths[i].index < strengths[j].index
		}
		return strengths[i].strength < strengths[j].strength
	})

	// Собираем индексы первых k строк
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = strengths[i].index
	}

	return result
}

func main() {
	// Пример матрицы
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
	}
	k := 3

	// Вызов функции и вывод результата
	result := kWeakestRows(mat, k)
	fmt.Println(result) // Ожидаемый вывод: [2, 0, 3]
}
