package _8_count_and_say

import "strconv"

func countAndSayV1(n int) string {
	/*
		Method: Recursion
		Time complexity: O(n*m), где n - входное число, m - длина строки, полученной на предыдущем шаге.
		Space complexity: O(n), так как мы храним только две строки в памяти: текущую и предыдущую.

		Time complexity
		Временная сложность этого алгоритма - O(n*m), где n - входное число, m - длина строки, полученной на предыдущем шаге.
		Это связано с тем, что для каждого символа в строке мы проверяем, является ли он частью той же группы, что и предыдущий символ.

		Space complexity
		Пространственная сложность - O(n), так как мы храним только две строки в памяти: текущую и предыдущую.
	*/

	// Базовый случай
	if n == 1 {
		return "1"
	}

	// Рекурсивный случай
	prev := countAndSayV1(n - 1)
	result := ""
	count := 1

	for i := 1; i < len(prev); i++ {
		if prev[i-1] == prev[i] {
			count++
		} else {
			result += strconv.Itoa(count) + string(prev[i-1])
			count = 1
		}
	}

	result += strconv.Itoa(count) + string(prev[len(prev)-1])

	return result
}
