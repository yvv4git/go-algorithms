//go:build ignore

package main

import "fmt"

/*
Approach: Модификация исходного массива.
- Основной алгоритмический подход — линейный проход (linear scan).
- Трюк — in-place marking (маркировка по индексу с помощью изменения знака).
Для каждого числа по его абсолютному значению вычисляем индекс (num-1) и меняем знак по этому индексу на отрицательный.
Если по индексу уже стоит отрицательное число, значит это число встречается второй раз — добавляем его в результат.

Time complexity: O(n), где n — дл��на массива, так как мы проходим по массиву только один раз.
Space complexity: O(1) — дополнительная память не используется, кроме памяти под результат.
*/

// findDuplicates возвращает все элементы, которые встречаются дважды в массиве nums.
// Решение работает за O(n) времени и использует O(1) дополнительной памяти (не считая результата).
func findDuplicates(nums []int) []int {
	res := []int{}

	for _, num := range nums {
		idx := abs(num) - 1
		if nums[idx] < 0 {
			res = append(res, abs(num))
		} else {
			nums[idx] = -nums[idx]
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	// Пример 1: nums = [4,3,2,7,8,2,3,1]
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("Пример 1 (ожидается [2,3]):", findDuplicates(nums1))

	// Пример 2: nums = [1,1,2]
	nums2 := []int{1, 1, 2}
	fmt.Println("Пример 2 (ожидается [1]):", findDuplicates(nums2))
}
