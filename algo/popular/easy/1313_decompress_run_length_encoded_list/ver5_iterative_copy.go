//go:build ignore

package main

import "fmt"

func decompressRLElist(nums []int) []int {
	/*
		APPROACH: Iterative (с шагом 2)
		1. Проходим по nums с шагом 2 (i увеличивается на 2)
		2. Частота повторения freq = nums[i]
		3. Значение val = nums[i+1]
		4. Создаём новый массив newPart с длиной freq и заполняем его val
		5. Добавляем newPart в result
		6. Повторяем шаги 2-5 для всех пар [freq, val] в nums

		TIME COMPLEXITY: O(n)
		- Мы проходим по nums один раз, выполняя константные операции на каждом шаге.

		SPACE COMPLEXITY: O(n)
		- Мы используем память для хранения результирующего списка.
		- Результирующий список содержит все элементы nums, поэтому его размер равен len(nums)/2.
	*/
	result := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
		freq, val := nums[i], nums[i+1]
		newPart := make([]int, freq)
		for j := range newPart {
			newPart[j] = val
		}
		result = append(result, newPart...)
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
