package main

import "fmt"

// Функция для поиска первого и последнего вхождения элемента в отсортированном массиве.
func searchRange(nums []int, target int) []int {
	// Инициализируем начальный и конечный индексы.
	start, end := -1, -1

	// Проверяем, что массив не пустой.
	if len(nums) == 0 {
		return []int{start, end}
	}

	// Бинарный поиск для первого вхождения.
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// Проверяем, что искомый элемент найден.
	if nums[left] != target {
		return []int{start, end}
	}

	// Запоминаем начальный индекс.
	start = left

	// Бинарный поиск для последнего вхождения.
	right = len(nums) - 1
	for left < right {
		mid := left + (right-left)/2 + 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}

	// Запоминаем конечный индекс.
	end = right

	// Возвращаем диапазон.
	return []int{start, end}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	result := searchRange(nums, target)
	fmt.Println(result) // Должно вывести [3, 4]
}
