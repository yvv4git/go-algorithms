package main

func sumOddLengthSubarraysV3(arr []int) int {
	n := len(arr)
	totalSum := 0

	for i := 0; i < n; i++ {
		// Количество подмассивов нечетной длины, включающих arr[i]
		count := ((i+1)*(n-i) + 1) / 2
		totalSum += arr[i] * count
	}

	return totalSum
}
