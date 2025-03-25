//go:build ignore

package main

import "fmt"

// Нет переаллокаций (append не используется).
// Эффективнее по памяти (особенно для больших списков).

func decompressRLElist(nums []int) []int {
	/*
		APPROACH: Iterative (оптимизированная)
		1. Вычисляем общее количество элементов
		2. Создаём слайс нужной длины
		3. Заполняем result
		4. Возвращаем result

		TIME COMPLEXITY: O(n)
		- Мы проходим по nums один раз, выполняя константные операции на каждом шаге.

		SPACE COMPLEXITY: O(n)
		- Мы используем память для хранения результирующего списка.
		- Результирующий список содержит все элементы nums, поэтому его размер равен len(nums)/2.
	*/
	// 1. Вычисляем общее количество элементов
	total := 0
	for i := 0; i < len(nums); i += 2 {
		total += nums[i]
	}

	// 2. Создаём слайс нужной длины
	result := make([]int, total)
	pos := 0 // Текущая позиция в result

	// 3. Заполняем result
	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		for j := 0; j < freq; j++ {
			result[pos] = val
			pos++
		}
	}

	return result
}

func main() {
	// Пример 1
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums1)) // [2 4 4 4]

	// Пример 2
	nums2 := []int{2, 1, 3, 5}
	fmt.Println(decompressRLElist(nums2)) // [1 1 5 5 5]
}
