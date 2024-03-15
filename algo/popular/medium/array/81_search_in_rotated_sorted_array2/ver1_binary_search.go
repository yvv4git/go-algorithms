package main

import "fmt"

// Функция для бинарного поиска в отсортированном и циклически сдвинутом массиве
func search(nums []int, target int) bool {
	/*
		METHOD: Binary search
		В этом решении мы используем бинарный поиск, но учитываем дополнительный случай,
		когда левая и правая границы могут совпадать, и при этом они равны целевому элементу.
		В этом случае мы уменьшаем диапазон поиска, иначе мы продолжаем бинарный поиск как обычно.
		Если целевого элемента нет в массиве, функция вернет false.

		TIME COMPLEXITY: O(log n) в среднем и в худшем случае, где n - количество элементов в массиве.

		SPACE COMPLEXITY: O(1), так как мы используем только несколько переменных, не связанных с размером входных данных.
	*/
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		// Если левая и правая границы равны, то уменьшаем диапазон поиска
		if nums[left] == nums[mid] && nums[right] == nums[mid] {
			left++
			right--
		} else if nums[left] <= nums[mid] { // Левая часть отсортирована
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1 // Ищем в левой половине
			} else {
				left = mid + 1 // Ищем в правой половине
			}
		} else { // Правая часть отсортирована
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1 // Ищем в правой половине
			} else {
				right = mid - 1 // Ищем в левой половине
			}
		}
	}

	return false // Если элемент не найден
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target)) // Вывод: true

	target = 3
	fmt.Println(search(nums, target)) // Вывод: false
}
