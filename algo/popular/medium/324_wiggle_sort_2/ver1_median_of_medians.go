package main

import (
	"fmt"
	"sort"
)

// wiggleSort осуществляет "Wiggle Sort II" для массива nums.
func wiggleSort(nums []int) {
	/*
		METHOD: Median of medians
		Для решения задачи "Wiggle Sort II" на языке Go мы можем использовать подход, основанный на сортировке и разделении массива на две части.
		Этот подход называется "median of medians" или "медиана медиан", который позволяет найти медиану массива за линейное время.
		Затем мы используем эту медиану для разделения массива на две части и чередования элементов.

		Объяснение подхода:
		1. Нахождение медианы: Сначала мы сортируем массив и находим медиану. Это делается за время O(n log n) из-за сортировки.
		2. Виртуальные индексы: Мы используем функцию virtIndex для получения виртуального индекса, который помогает нам чередовать элементы правильным образом.
		3. Чередование элементов: Мы проходим по массиву и чередуем элементы вокруг медианы, используя виртуальные индексы.

		TIME COMPLEXITY: O(n log n) из-за сортировки массива для нахождения медианы.

		SPACE COMPLEXITY: O(1), так как мы модифицируем исходный массив без использования дополнительной памяти, кроме нескольких переменных.
	*/
	n := len(nums)
	if n <= 1 {
		return
	}

	// Находим медиану массива
	median := findMedian(nums)

	// Индексы для чередования элементов
	left, i, right := 0, 0, n-1

	// Функция для получения виртуального индекса
	virtIndex := func(idx int) int {
		return (1 + 2*idx) % (n | 1)
	}

	// Чередуем элементы вокруг медианы
	for i <= right {
		if nums[virtIndex(i)] > median {
			nums[virtIndex(left)], nums[virtIndex(i)] = nums[virtIndex(i)], nums[virtIndex(left)]
			left++
			i++
		} else if nums[virtIndex(i)] < median {
			nums[virtIndex(i)], nums[virtIndex(right)] = nums[virtIndex(right)], nums[virtIndex(i)]
			right--
		} else {
			i++
		}
	}
}

// findMedian находит медиану массива nums
func findMedian(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n/2]
}

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums) // Ожидаемый вывод: [1, 6, 1, 5, 1, 4]
}
