package main

import (
	"fmt"
)

func createTargetArray(nums []int, index []int) []int {
	/*
		APPROACH: Linear scan
		1. Создать пустой массив target длиной len(nums)
		2. Пройтись по index и nums в цикле
		3. Вставляем элемент nums[i] в позицию index[i] в target
		4. Возвращаем target

		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)
	*/
	target := make([]int, 0, len(nums)) // Инициализация с емкостью для оптимизации
	for i, num := range nums {
		pos := index[i]
		// Вставка элемента num на позицию pos
		target = append(target[:pos], append([]int{num}, target[pos:]...)...)
	}

	return target
}

func main() {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}
	result := createTargetArray(nums, index)
	fmt.Println(result) // Вывод: [0, 4, 1, 3, 2]
}
