package main

import "fmt"

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
func miniMaxSum(arr []int32) []int32 {
	// Write your code here
	if len(arr) != 5 {
		return []int32{}
	}

	min := int32(0)
	max := int32(0)

	bubbleSort(arr) // sort before

	length := len(arr)
	for i := 0; i < 4; i++ {
		min += arr[i]
		max += arr[length-i-1]
	}

	fmt.Println(min, max)
	return []int32{min, max}
}

// Пузырьковая сортировка, нужна, чтобы отсортировать исходный массив.
func bubbleSort(arr []int32) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
