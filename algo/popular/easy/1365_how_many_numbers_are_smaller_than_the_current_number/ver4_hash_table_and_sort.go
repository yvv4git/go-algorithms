//go:build ignore

package main

import (
	"sort"
)

func smallerNumbersThanCurrent(nums []int) []int {
	/*
		METHOD: Hash table and sort
		1. Сортируем массив
		2. Создаем хэш-таблицу для хранения первого вхождения каждого числа

		TIME COMPLEXITY: O(nlogn) + O(n) = O(nlogn)
		SPACE COMPLEXITY: O(n)
	*/
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums) // Сортируем копию массива

	// Создаем хэш-таблицу для хранения первого вхождения каждого числа
	hash := make(map[int]int)
	for i, num := range sortedNums {
		if _, exists := hash[num]; !exists {
			hash[num] = i
		}
	}

	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = hash[num]
	}
	return result
}
