package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	/*
		METHOD: Custom sorting
		Time complexity: O(n log n), где n - количество чисел в массиве
		Space complexity: O(n), так как для сортировки требуется дополнительное пространство для хранения временного массива.
	*/
	// Преобразуем числа в строки
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}

	// Сортируем строки с использованием кастомной функции сравнения
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})

	// Если первый элемент равен "0", то и все остальные элементы будут "0",
	// поэтому мы можем вернуть "0" сразу
	if strs[0] == "0" {
		return "0"
	}

	// Объединяем строки в одну
	return strings.Join(strs, "")
}
