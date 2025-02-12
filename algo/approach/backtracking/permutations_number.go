//go:build ignore

package main

import (
	"fmt"
)

func backtrack(nums []int, path []int, used []bool, result *[][]int) {
	// Когда длина текущего пути равна длине массива, добавляем его в результат
	if len(path) == len(nums) {
		// Создаем копию пути, чтобы избежать изменений при откате
		perm := make([]int, len(path))
		copy(perm, path)
		*result = append(*result, perm)
		return
	}

	// Перебираем все элементы массива
	for i := 0; i < len(nums); i++ {
		// Пропускаем уже использованные элементы
		if used[i] {
			continue
		}

		// Добавляем текущий элемент в путь
		used[i] = true
		path = append(path, nums[i])

		// Рекурсивный вызов для продолжения построения пути
		backtrack(nums, path, used, result)

		// Откат (удаление последнего элемента из пути и отметка его как неиспользованного)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func permute(nums []int) [][]int {
	var result [][]int
	used := make([]bool, len(nums))
	backtrack(nums, []int{}, used, &result)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)

	// Выводим результат
	fmt.Println("Перестановки:")
	for _, perm := range result {
		fmt.Println(perm)
	}
}
