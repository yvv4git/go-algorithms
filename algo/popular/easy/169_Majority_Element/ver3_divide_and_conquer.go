//go:build ignore

package main

import "fmt"

/*
APPROACH:  Divide and Conquer
Рекурсивно делим массив на две части, ищем majority element в каждой части.
Если оба совпадают — возвращаем его.
Иначе считаем количество каждого кандидата на текущем отрезке и возвращаем того,
кто встречается чаще.

TIME COMPLEXITY: "O(n log n), так как на каждом уровне рекурсии делим массив пополам,
а при слиянии считаем количество кандидато��."

SPACE COMPLEXITY: "O(log n), так как глубина рекурсии пропорциональна log n."
*/
func majorityElement(nums []int) int {
	var majority func(lo, hi int) int

	majority = func(lo, hi int) int {
		if lo == hi {
			return nums[lo]
		}

		mid := (lo + hi) / 2
		left := majority(lo, mid)
		right := majority(mid+1, hi)
		if left == right {
			return left
		}
		leftCount, rightCount := 0, 0

		for i := lo; i <= hi; i++ {
			if nums[i] == left {
				leftCount++
			} else if nums[i] == right {
				rightCount++
			}
		}

		if leftCount > rightCount {
			return left
		}

		return right
	}

	return majority(0, len(nums)-1)
}

func main() {
	examples := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}
	for _, nums := range examples {
		res := majorityElement(nums)
		fmt.Printf("Input: %v\n", nums)
		fmt.Printf("Majority element: %d\n\n", res)
	}
}
