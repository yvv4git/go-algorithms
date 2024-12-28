package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	/*
		METHOD: Prefix Sum
		1. Используем массив префиксных сумм для быстрого вычисления суммы любого подмассива.
		2. Перебираем все подмассивы нечетной длины (1, 3, 5, ...).
		3. Для каждого подмассива вычисляем его сумму через разность префиксных сумм и добавляем к общей сумме.

		TIME COMPLEXITY: O(n^2)
		1. Вычисление префиксных сумм занимает O(n).
		2. Перебор всех подмассивов нечетной длины требует O(n^2), так как для каждой длины (n/2 вариантов)
		   мы перебираем начальные индексы (до n вариантов).
		3. Итоговая сложность: O(n^2).

		SPACE COMPLEXITY: O(n)
		1. Используем дополнительный массив для хранения префиксных сумм размером n+1.
		2. Итоговая пространственная сложность: O(n).
	*/
	n := len(arr)
	prefixSum := make([]int, n+1)

	// Вычисляем префиксные суммы
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + arr[i]
	}

	totalSum := 0

	// Перебираем все подмассивы нечетной длины
	for length := 1; length <= n; length += 2 {
		for i := 0; i <= n-length; i++ {
			// Сумма подмассива от i до i+length-1
			totalSum += prefixSum[i+length] - prefixSum[i]
		}
	}

	return totalSum
}

func main() {
	arr := []int{1, 4, 2, 5, 3}
	result := sumOddLengthSubarrays(arr)
	fmt.Println(result) // Вывод: 58
}
