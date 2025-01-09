package main

func numIdenticalPairsV2(nums []int) int {
	/*
		METHOD: Map
		Создаём счётчик для подсчёта встречающихся чисел,
		и при каждом новом вхождении числа добавляем к общему результату количество его предыдущих появлений - таким образом получаем все возможные пары с этим числом.

		TIME COMPLEXITY: O(n), где n - длина массива `nums`.
		SPACE COMPLEXITY: O(n), так как используется карта `count` для хранения количества вхождений чисел.
	*/
	count := make(map[int]int)
	result := 0

	for _, num := range nums {
		result += count[num]
		count[num]++
	}

	return result
}
