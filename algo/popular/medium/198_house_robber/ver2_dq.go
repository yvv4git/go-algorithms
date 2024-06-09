package main

import "fmt"

func robV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robHelper(nums, 0, len(nums)-2), robHelper(nums, 1, len(nums)-1))
}

func robHelper(nums []int, start int, end int) int {
	prev, curr := 0, 0
	for i := start; i <= end; i++ {
		temp := max(curr, prev+nums[i])
		prev = curr
		curr = temp
	}
	return curr
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

func main() {
	// Пример использования
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums)) // Вывод: 4
}
