package main

import "fmt"

func duplicateZeros(arr []int) {
	/*
		METHOD:

		TIME COMPLEXITY:

		SPACE COMPLEXITY:
	*/
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}

func main() {
	// Пример использования функции duplicateZeros
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Println("Original array:", arr)
	duplicateZeros(arr)
	fmt.Println("Array after duplicating zeros:", arr)
}
