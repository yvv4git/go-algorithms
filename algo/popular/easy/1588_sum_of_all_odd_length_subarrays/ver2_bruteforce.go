package main

func sumOddLengthSubarraysV2(arr []int) int {
	n := len(arr)
	totalSum := 0

	// Перебираем все подмассивы нечетной длины
	for length := 1; length <= n; length += 2 {
		for i := 0; i <= n-length; i++ {
			// Вычисляем сумму текущего подмассива
			sum := 0
			for j := i; j < i+length; j++ {
				sum += arr[j]
			}
			totalSum += sum
		}
	}

	return totalSum
}
