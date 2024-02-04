package main

import (
	"sort"
)

func majorityElementSort(nums []int) int {
	/*
		METHOD: Sorting
		Идея состоит в том, что если отсортировать элементы, то тот самый majority элемент будет посередине.

		Time complexity : O(nlgn)
		Space complexity : O(1) or O(n) - все зависит от алгоритма сортировки.
	*/
	count := len(nums)

	if count == 0 {
		return 0
	}

	sort.Ints(nums)
	idx := count / 2

	return nums[idx]
}
