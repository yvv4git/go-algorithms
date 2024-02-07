package main

import "fmt"

// Функция для поиска индекса опорного элемента
func findPivot(nums []int) int {
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			return i - 1
		}
	}
	return -1
}

// Функция для поиска индекса меньшего элемента
func findSmaller(nums []int, pivot int) int {
	for i := len(nums) - 1; i > pivot; i-- {
		if nums[i] > nums[pivot] {
			return i
		}
	}
	return -1
}

// Функция для переворота элементов справа от опорного
func reverse(nums []int, start int) {
	end := len(nums) - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// Функция для нахождения следующей перестановки
func nextPermutation(nums []int) {
	/*
		METHOD: Next Permutation
		TIME COMPLEXITY: O(n), где n - количество элементов в массиве.
		Это связано с двумя циклами, которые проходят по всему массиву: один для поиска опорного элемента
		и один для поиска меньшего элемента, а также для переворота элементов.
		SPACE COMPLEXITY: O(1), так как мы используем только несколько дополнительных переменных
		и не используем дополнительную память, зависящую от размера входных данных.
	*/
	// Шаг 1: Находим опорный элемент
	pivot := findPivot(nums)
	if pivot != -1 {
		// Шаг 2: Находим меньший элемент
		smaller := findSmaller(nums, pivot)
		// Шаг 3: Меняем местами опорный и меньший элементы
		nums[pivot], nums[smaller] = nums[smaller], nums[pivot]
	}
	// Шаг 4: Переворачиваем элементы справа от опорного
	reverse(nums, pivot+1)
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // Вывод: [1 3 2]
}
