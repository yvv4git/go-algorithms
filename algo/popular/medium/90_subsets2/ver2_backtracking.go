package main

import "sort"

// Функция для получения всех подмножеств с возможными дубликатами
func subsetsWithDupV2(nums []int) [][]int {
	/*
		METHOD: Backtracking

		TIME COMPLEXITY: O(N * 2^N) + O(NlogN) -> O(N * 2^N)

		SPACE COMPLEXITY: O(2^N) - Output space
	*/
	// Сортируем исходный массив
	sort.Ints(nums)
	// Инициализируем результирующий массив
	res := [][]int{}

	// Проходим по всем длинам подмножеств, которые мы хотим получить
	for i := 0; i <= len(nums); i++ {
		// Ищем все подмножества длины i
		findSubset(nums, 0, i, []int{}, &res)
	}
	// Возвращаем результат
	return res
}

// Рекурсивная функция для поиска подмножеств
func findSubset(nums []int, idx, l int, list []int, res *[][]int) {
	// Если длина текущего списка равна l, то добавляем его в результат
	if len(list) == l {
		temp := make([]int, l)
		copy(temp, list)
		(*res) = append(*res, temp)
		return
	}

	// Проходим по всем элементам, начиная с idx
	for i := idx; i < len(nums); i++ {
		// Если это первый элемент или он не равен предыдущему, то добавляем его в список
		if i == idx || nums[i] != nums[i-1] {
			findSubset(nums, i+1, l, append(list, nums[i]), res)
		}
	}
}
