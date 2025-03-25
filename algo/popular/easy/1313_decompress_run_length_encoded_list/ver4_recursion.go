//go:build ignore

package main

import "fmt"

func decompressRLElist(nums []int) []int {
	/*
		APPROACH: Recursion
		1. Проверяем, что nums не пустой
		2. Частота повторения freq = nums[0]
		3. Значение val = nums[1]
		4. Создаём новый массив current с длиной freq и заполняем его val
		5. Добавляем current в результирующий список
		6. Повторяем шаги 2-5 для всех пар [freq, val] в nums
		7. Возвращаем результирующий список

		TIME COMPLEXITY: O(n)
		- Мы проходим по nums один раз, выполняя константные операции на каждом шаге.

		SPACE COMPLEXITY: O(n)
		- Мы используем память для хранения результирующего списка.
		- Результирующий список содержит все элементы nums, поэтому его размер равен len(nums)/2.
	*/
	if len(nums) == 0 {
		return []int{}
	}

	freq, val := nums[0], nums[1]
	current := make([]int, freq)
	for i := 0; i < freq; i++ {
		current[i] = val
	}

	return append(current, decompressRLElist(nums[2:])...)
}

func main() {
	// Пример 1
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(decompressRLElist(nums1)) // [2 4 4 4]

	// Пример 2
	nums2 := []int{2, 1, 3, 5}
	fmt.Println(decompressRLElist(nums2)) // [1 1 5 5 5]
}
