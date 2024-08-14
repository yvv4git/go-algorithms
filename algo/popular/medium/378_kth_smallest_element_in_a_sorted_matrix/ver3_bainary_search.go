package main

import "sort"

func kthSmallestV3(matrix [][]int, k int) int {
	/*
		METHOD: Binary Search
		DESCRIPTION: Используется двоичный поиск по значениям матрицы для нахождения k-го наименьшего элемента.
		Двоичный поиск используется в данной задаче из-за отсортированности матрицы (по строкам и столбцам),
		что позволяет эффективно сужать область поиска за логарифмическое время, обеспечивая быстрое нахождение k-го наименьшего элемента.

		TIME COMPLEXITY: O(n log n log(max-min))
		DESCRIPTION: Временная сложность определяется двоичным поиском, где n — размер матрицы,
		max и min — максимальный и минимальный элементы матрицы соответственно.
		Каждая итерация двоичного поиска требует O(n log n) времени для подсчета элементов с использованием функции sort.Search,
		а общее количество итераций двоичного поиска составляет O(log(max-min)).

		SPACE COMPLEXITY: O(1)
		DESCRIPTION: Используется только константное количество дополнительной памяти,
		поэтому пространственная сложность равна O(1).
	*/
	// Инициализируем границы поиска: minVal - минимальный элемент, maxVal - максимальный элемент
	minVal, maxVal := matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]

	// Начинаем двоичный поиск
	for minVal < maxVal {
		// Вычисляем среднее значение между minVal и maxVal
		midVal := minVal + (maxVal-minVal)/2

		// Если количество элементов, меньших или равных midVal, меньше k, сдвигаем minVal вправо
		if countElementsLessOrEqual(matrix, midVal) < k {
			minVal = midVal + 1
		} else {
			// В противном случае сдвигаем maxVal влево
			maxVal = midVal
		}
	}

	// Возвращаем k-й наименьший элемент
	return minVal
}

func countElementsLessOrEqual(matrix [][]int, target int) int {
	// Инициализируем счетчик для количества элементов, меньших или равных target
	totalCount := 0

	// Проходим по каждой строке матрицы
	for _, row := range matrix {
		// Используем sort.Search для нахождения количества элементов, меньших или равных target в текущей строке
		totalCount = totalCount + sort.Search(len(row), func(i int) bool {
			return row[i] > target
		})
	}

	// Возвращаем общее количество элементов, меньших или равных target
	return totalCount
}
