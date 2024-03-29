package main

import "sort"

// Функция containsDuplicateV2 проверяет, содержит ли массив дубликаты.
// Она использует сортировку для этого.
func containsDuplicateV2(nums []int) bool {
	/*
		METHOD: Sort
		TIME COMPLEXITY: O(n log n)
		SPACE COMPLEXITY: O(1)

		В этом коде мы сначала сортируем массив, а затем проходим по нему и проверяем, равны ли соседние элементы.
		Если они равны, то мы нашли дубликат, и функция возвращает true.
		Если мы дошли до конца массива и не нашли дубликатов, то функция возвращает false.
	*/

	// Сортируем массив - O(n log n).
	sort.Ints(nums)

	// Проходим по массиву O(n).
	for i := 0; i < len(nums)-1; i++ {
		// Если текущий элемент равен следующему, то в массиве есть дубликаты.
		if nums[i] == nums[i+1] {
			return true
		}
	}

	// Если мы дошли до этой точки, то в массиве нет дубликатов.
	return false
}
