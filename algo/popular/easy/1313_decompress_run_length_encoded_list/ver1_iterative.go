package main

import "fmt"

func decompressRLElist(nums []int) []int {
	/*
		APPROACH: Iterative (с шагом 2)
		1. Проходим по nums с шагом 2 (i увеличивается на 2)
		2. Частота повторения freq = nums[i]
		3. Значение val = nums[i+1]
		4. Добавляем val в result freq раз

		TIME COMPLEXITY: O(n)
		- Мы проходим по nums один раз, выполняя константные операции на каждом шаге.

		SPACE COMPLEXITY: O(n)
		- Мы используем память для хранения результирующего списка.
		- Результирующий список содержит все элементы nums, поэтому его размер равен len(nums)/2.
	*/
	var result []int // Создаём пустой слайс для результата

	// Проходим по nums с шагом 2 (i увеличивается на 2)
	for i := 0; i < len(nums); i += 2 {
		freq := nums[i]  // Частота повторения
		val := nums[i+1] // Значение, которое нужно повторить

		// Добавляем val в result freq раз
		for j := 0; j < freq; j++ {
			result = append(result, val)
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
