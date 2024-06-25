package main

func findDuplicate(nums []int) int {
	/*
		METHOD: Быстрый и медленный указатели (Floyd's Tortoise and Hare algorithm)

		TIME COMPLEXITY: O(n), так как алгоритм выполняет два прохода по массиву, каждый из которых занимает линейное время.

		SPACE COMPLEXITY: O(1), так алгоритм использует фиксированное количество переменных, не зависящее от размера входного массива.
	*/

	// Шаг 1: Найти точку встречи быстрого и медленного указателей
	slow := nums[0]
	fast := nums[0]

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	// Шаг 2: Найти начало цикла
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
