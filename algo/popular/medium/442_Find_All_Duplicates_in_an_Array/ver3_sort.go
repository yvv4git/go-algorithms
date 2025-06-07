package main

import (
	"fmt"
	"sort"
)

/*
Approach: Сортировка массива.
После сортировки все дубликаты будут идти подряд, и их можно найти одним проходом.

Time complexity: O(n log n) — из-за сортировки.
Space complexity: O(1) — если сортировка in-place, иначе зависит от реализации sort.
*/

func findDuplicatesSort(nums []int) []int {
	res := []int{}
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			// Чтобы не добавить один и тот же дубликат несколько раз
			if len(res) == 0 || res[len(res)-1] != nums[i] {
				res = append(res, nums[i])
			}
		}
	}

	return res
}

func main() {
	// Пример 1: nums = [4,3,2,7,8,2,3,1]
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("Пример 1 (ожидается [2,3]):", findDuplicatesSort(nums1))

	// Пример 2: nums = [1,1,2]
	nums2 := []int{1, 1, 2}
	fmt.Println("Пример 2 (ожидается [1]):", findDuplicatesSort(nums2))
}
