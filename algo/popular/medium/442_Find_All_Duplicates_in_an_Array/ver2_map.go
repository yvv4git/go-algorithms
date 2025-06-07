//go:build ignore

package main

import "fmt"

/*
Approach: Использование map/set для поиска дубликатов.
Для каждого числа проверяем, встречалось ли оно ранее (через map). Если да — добавляем в результат.

Time complexity: O(n)
Space complexity: O(n)
*/

func findDuplicatesMap(nums []int) []int {
	seen := make(map[int]bool)
	res := []int{}
	for _, num := range nums {
		if seen[num] {
			res = append(res, num)
		} else {
			seen[num] = true
		}
	}
	return res
}

func main() {
	// Пример 1: nums = [4,3,2,7,8,2,3,1]
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("Пример 1 (ожидается [2,3]):", findDuplicatesMap(nums1))

	// Пример 2: nums = [1,1,2]
	nums2 := []int{1, 1, 2}
	fmt.Println("Пример 2 (ожидается [1]):", findDuplicatesMap(nums2))
}
