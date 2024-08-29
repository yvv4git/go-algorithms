package main

import (
	"math"
	"sort"
)

// maxSumSubmatrix находит максимальную сумму подматрицы, не превышающую k.
// Method: Префиксные суммы и бинарный поиск.
// Time Complexity: O(cols^2 * rows * log(rows))
// Space Complexity: O(rows)
func maxSumSubmatrix(matrix [][]int, k int) int {
	/*
		METHOD:
		- Используется метод префиксных сумм и бинарный поиск.
		- Для каждой пары столбцов (left, right) вычисляются суммы строк между этими столбцами.
		- Для каждого набора сумм строк вызывается функция maxSumSubarray, которая находит максимальную сумму подмассива, не превышающую k.

		TIME COMPLEXITY:
		- O(cols^2 * rows * log(rows)):
			- Проход по всем парам столбцов: O(cols^2).
			- Для каждой пары столбцов вычисление сумм строк: O(rows).
			- Для каждого набора сумм строк вызов maxSumSubarray: O(rows * log(rows)).

		SPACE COMPLEXITY:
		- O(rows):
			- Для хранения сумм строк между столбцами используется массив temp размером rows.
	*/
	rows := len(matrix)
	cols := len(matrix[0])
	maxSum := math.MinInt32

	// Проходим по всем возможным парам столбцов
	for left := 0; left < cols; left++ {
		// temp будет хранить суммы строк между left и right столбцами
		temp := make([]int, rows)
		for right := left; right < cols; right++ {
			// Обновляем суммы строк для текущей пары столбцов
			for i := 0; i < rows; i++ {
				temp[i] += matrix[i][right]
			}
			// Находим максимальную сумму подмассива, не превышающую k
			maxSum = max(maxSum, maxSumSubarray(temp, k))
			if maxSum == k {
				return k // Ранний выход, если найдена точная сумма k
			}
		}
	}
	return maxSum
}

// maxSumSubarray находит максимальную сумму подмассива, не превышающую k.
// Method: Использование TreeSet для эффективного поиска.
// Time Complexity: O(rows * log(rows))
// Space Complexity: O(rows)
func maxSumSubarray(arr []int, k int) int {
	set := make([]int, 0)
	set = append(set, 0)
	currentSum := 0
	maxSum := math.MinInt32

	for _, num := range arr {
		currentSum += num
		// Ищем ближайшее число в set, которое currentSum - x <= k
		index := sort.Search(len(set), func(i int) bool {
			return set[i] >= currentSum-k
		})
		if index < len(set) {
			maxSum = max(maxSum, currentSum-set[index])
		}
		// Вставляем currentSum в отсортированный массив set
		index = sort.Search(len(set), func(i int) bool {
			return set[i] >= currentSum
		})
		set = append(set, 0)
		copy(set[index+1:], set[index:])
		set[index] = currentSum
	}
	return maxSum
}

// max возвращает максимальное значение из двух целых чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
