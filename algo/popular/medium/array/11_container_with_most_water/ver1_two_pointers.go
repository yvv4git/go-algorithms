package main

import "fmt"

// Функция для нахождения максимального объема воды
func maxArea(height []int) int {
	/*
		METHOD: Two pointers
		TIME COMPLEXITY: O(n), где n - количество элементов в массиве. Это обусловлено тем, что мы проходим по массиву всего один раз.
		SPACE COMPLEXITY: O(1), так как мы используем только несколько переменных для хранения индексов и максимального объема.
	*/
	// Инициализация переменных для указателей и максимального объема
	left, right := 0, len(height)-1
	maxArea := 0

	// Цикл, в котором мы будем двигать указатели
	for left < right {
		// Вычисление текущего объема воды и обновление максимального объема
		width := right - left
		maxArea = max(maxArea, min(height[left], height[right])*width)

		// Перемещение указателей
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	// Возвращение максимального объема
	return maxArea
}

// Функция для нахождения минимального значения из двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Функция для нахождения максимального значения из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример использования функции
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height)) // Вывод: 49
}
