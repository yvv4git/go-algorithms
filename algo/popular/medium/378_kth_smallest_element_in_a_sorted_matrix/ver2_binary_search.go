package main

import "sort"

func kthSmallestV2(matrix [][]int, k int) int {
	/*
		METHOD: Binary Search
		DESCRIPTION: Используется двоичный поиск по значениям матрицы для нахождения k-го наименьшего элемента.

		TIME COMPLEXITY: O(n log n log(max-min))
		DESCRIPTION: Временная сложность определяется двоичным поиском, где n — размер матрицы,
		max и min — максимальный и минимальный элементы матрицы соответственно.
		Каждая итерация двоичного поиска требует O(n log n) времени для подсчета элементов с использованием функции sort.Search,
		а общее количество итераций двоичного поиска составляет O(log(max-min)).

		SPACE COMPLEXITY: O(1)
		DESCRIPTION: Используется только константное количество дополнительной памяти,
		поэтому пространственная сложность равна O(1).
	*/
	// Инициализируем границы поиска: i - минимальный элемент, j - максимальный элемент
	i, j := matrix[0][0], matrix[len(matrix)-1][len(matrix[0])-1]

	// Начинаем двоичный поиск
	for i < j {
		// Вычисляем среднее значение между i и j
		mid := i + (j-i)/2

		// Если количество элементов, меньших или равных mid, меньше k, сдвигаем i вправо
		if smallerCount(matrix, mid) < k {
			i = mid + 1
		} else {
			// В противном случае сдвигаем j влево
			j = mid
		}
	}

	// Возвращаем k-й наименьший элемент
	return i
}

func smallerCount(matrix [][]int, k int) int {
	// Инициализируем счетчик для количества элементов, меньших или равных k
	ret := 0

	// Проходим по каждой строке матрицы
	for _, v := range matrix {
		// Используем sort.Search для нахождения количества элементов, меньших или равных k в текущей строке
		ret = ret + sort.Search(len(v), func(i int) bool {
			return v[i] > k
		})
	}

	// Возвращаем общее количество элементов, меньших или равных k
	return ret
}