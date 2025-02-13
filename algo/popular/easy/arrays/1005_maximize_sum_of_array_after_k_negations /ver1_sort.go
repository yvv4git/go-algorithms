//go:build ignore

package main

import (
	"fmt"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	/*
		METHOD: Iterate after sort
		Для решения этой задачи мы будем использовать прием,
		основанный на сортировке и изменении знака минимальных элементов массива.

		TIME COMPLEXITY: O(n log n), где n - количество элементов в массиве, так как мы используем сортировку.

		SPACE COMPLEXITY: O(1), так как мы не используем дополнительное пространство, кроме для переменных и входных данных.
	*/
	// Сортируем массив
	sort.Ints(nums)

	// Проходимся по массиву и переворачиваем отрицательные элементы
	for i := 0; i < len(nums) && k > 0 && nums[i] < 0; i++ {
		nums[i] = -nums[i]
		k--
	}

	// Если осталось нечетное количество операций, переворачиваем минимальный элемент
	if k%2 == 1 {
		minIndex := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[minIndex] {
				minIndex = i
			}
		}
		nums[minIndex] = -nums[minIndex]
	}

	// Считаем сумму элементов
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	nums := []int{4, 2, 3}
	K := 1
	fmt.Println(largestSumAfterKNegations(nums, K)) // Выводит: 5
}
