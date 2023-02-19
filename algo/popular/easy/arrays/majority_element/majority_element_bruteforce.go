package main

func majorityElementBruteForce(nums []int) int {
	/*
		Method: Brute Force.
		Т.е.:
		0. Находим majority element value.
		1. Берем элемент.
		2. Проходим по всему списку элементов и считаем количество раз, когда выбранный элемент повторяется в списке.
		3. Проверяем на факт того, что количество появлений элемента в списке > majority element value.

		Time complexity: O(n^2)
		Space complexity: O(1)
	*/
	majorElem := len(nums) / 2

	for _, num := range nums {
		count := 0

		for _, num2 := range nums {
			if num2 == num {
				count++
			}

			if count > majorElem {
				return num
			}
		}

	}

	return 0
}
