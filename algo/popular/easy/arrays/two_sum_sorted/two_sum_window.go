package main

func twoSumWindow(nums []int, target int) []int {
	/*
		Method: Window
		Алгоритм простой и в тоже время эффективный.


		Время работы: O(n)
		Память: O(1)
	*/

	l, r := 0, len(nums)-1 // Изначальный размер окна

	for l < r {
		sum := nums[l] + nums[r]

		if sum == target {
			return []int{l, r}
		}

		if sum > target {
			r-- // Уменьшаем окно справа на 1
		} else {
			l++ // Уменьшаем окно слева на 1
		}
	}

	return []int{}
}
