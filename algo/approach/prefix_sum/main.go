package main

import "fmt"

// Функция для вычисления префиксной суммы массива
func prefixSum(arr []int) []int {
	/*
		METHOD: Prefix sum
		Префиксная сумма (prefix sum) - это массив, в котором каждый элемент является суммой всех предыдущих элементов исходного массива.
		В отличие от префиксной суммы, которая вычисляется для всего массива, суффиксная сумма - это массив,
		в котором каждый элемент является суммой всех последующих элементов исходного массива.

		TIME COMPLEXITY: O(n), где n - количество элементов в массиве.
		Это обусловлено тем, что мы проходим по каждому элементу массива ровно один раз.

		SPACE COMPLEXITY: O(n), так как мы создаем дополнительный массив для хранения префиксных сумм, который тоже содержит n элементов.
	*/
	n := len(arr)
	result := make([]int, n)

	// Первый элемент префиксной суммы равен первому элементу исходного массива
	result[0] = arr[0]

	// Вычисляем префиксную сумму для остальных элементов
	for i := 1; i < n; i++ {
		result[i] = result[i-1] + arr[i]
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	prefixSumArr := prefixSum(arr)

	fmt.Println("Префиксные суммы:", prefixSumArr)
}