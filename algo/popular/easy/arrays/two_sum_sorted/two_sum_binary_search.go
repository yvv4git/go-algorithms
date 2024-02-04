package main

func twoSumBinarySearch(nums []int, target int) []int {
	/*
		METHOD: Binary search.
		Используем знаменитый бинарный поиск для решения задачи.

		Есть одно но - массив должен быть отсортирован!

		Время работы: O(n log n)
		Память: O(1)
	*/

	for idx, num := range nums {
		want := target - num // Число, которое надо найти в массиве
		l, r := idx+1, len(nums)-1

		for l <= r {
			m := (l + r) / 2

			if nums[m] < want {
				l = m + 1
			} else if nums[m] > want {
				r = m - 1
			} else if nums[m] == want {
				return []int{idx, m}
			}
		}
	}

	return []int{}
}
