package main

func findMissingNumberV2(numbers []int, N int) int {
	/*
		METHOD: Math & Sum of the arithmetic sequence
		TIME COMPLEXITY: O(n), где n - количество элементов в последовательности numbers, так как мы проходим по всем элементам последовательности для вычисления суммы.
		SPACE COMPLEXITY: O(1), так как мы используем лишь несколько переменных для хранения промежуточных сумм.
	*/
	// Вычисляем сумму всех чисел от 1 до N
	totalSum := N * (N + 1) / 2

	// Вычисляем сумму чисел в данной последовательности
	sequenceSum := 0
	for _, number := range numbers {
		sequenceSum += number
	}

	// Возвращаем разность между суммой всех чисел и суммой последовательности,
	// которая и будет отсутствующим числом
	return totalSum - sequenceSum
}
