package main

import (
	"sort"
)

func threeSumClosestV2(nums []int, target int) int {
	/*
		METHOD: Binary search
		TIME COMPLEXITY: O(n^2). Это связано с тем, что бинарный поиск здесь не применим, так как требуется найти ближайшую к заданному числу сумму трех чисел, а не просто найти сумму, равную заданному числу.
		Бинарный поиск предназначен для поиска конкретного элемента в отсортированном массиве, а не для поиска ближайшего к заданному числу.
		SPACE COMPLEXITY: O(1), так как мы используем только несколько переменных для хранения промежуточных результатов, не зависящих от размера входных данных.
	*/
	sort.Ints(nums)
	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}

			if abs(target-sum) < abs(target-closestSum) {
				closestSum = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}

//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
//
//func main() {
//	nums := []int{-1, 2, 1, -4}
//	target := 1
//	fmt.Println(threeSumClosest(nums, target)) // Выводит: 2
//}
