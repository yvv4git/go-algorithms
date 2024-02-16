package main

// Функция для слияния двух отсортированных массивов
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(merged, right...)
		} else if len(right) == 0 {
			return append(merged, left...)
		} else if left[0] < right[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	return merged
}

// Функция для рекурсивного разделения массива и последующего слияния
func mergeSort(arr []int) []int {
	/*
		METHOD: Stack
		TIME COMPLEXITY:
		SPACE COMPLEXITY:

		В этом примере мы используем рекурсию для разделения массива на две половины до тех пор,
		пока каждая половина не будет состоять из одного элемента.
		Затем мы используем функцию merge для слияния отсортированных половинок обратно в отсортированный массив.

		Обратите внимание, что в этом примере стек не используется явно, так как рекурсия в Go реализуется неявно через стек.
		Однако, если бы мы хотели реализовать этот алгоритм без использования рекурсии,
		мы бы могли использовать стек вручную для отслеживания вызовов функции mergeSort.
	*/
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

//func main() {
//	arr := []int{10, 7, 8, 9, 1, 5}
//	fmt.Println("Original array:", arr)
//	sortedArr := mergeSort(arr)
//	fmt.Println("Sorted array:", sortedArr)
//}
