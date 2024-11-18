package main

import (
	"sort"
)

func repeatedNTimesV2(nums []int) int {
	/*
		METHOD: Sort
		TIME COMPLEXITY: O(n log(n))
		SPACE COMPLEXITY: O(1)
	*/
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// func main() {
//     nums := []int{1, 2, 3, 3}
//     result := repeatedNTimesV2(nums)
//     fmt.Println("Повторяющийся элемент:", result) // Вывод: Повторяющийся элемент: 3
// }
