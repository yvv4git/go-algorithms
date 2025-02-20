//go:build ignore

package main

import "fmt"

// Функция для нахождения максимальной суммы подмассива, который пересекает границу
func maxCrossingSum(arr []int, left, right, mid int) int {
	/*
		Задача: Найти максимальную сумму подмассива, который пересекает границу массива.

		Подход:
		1. Ищем максимальную сумму элементов в левой части, включая середину.
		2. Ищем максимальную сумму элементов в правой части, начиная с середины.
		3. Складываем эти суммы, включая средний элемент.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(n), где n — количество элементов в массиве.
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(1)
	*/
	leftSum := int(^uint(0) >> 1)  // минимальное возможное значение
	rightSum := int(^uint(0) >> 1) // минимальное возможное значение
	sum := 0
	for i := mid; i >= left; i-- {
		sum += arr[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += arr[i]
		if sum > rightSum {
			rightSum = sum
		}
	}

	return leftSum + rightSum
}

// Функция для нахождения максимальной суммы подмассива в массиве arr
func maxSubArraySum(arr []int, left, right int) int {
	/*
		Задача: Найти максимальную сумму подмассива в массиве arr.

		Подход:
		1. Разделяем массив на две части.
		2. Рекурсивно ищем максимальную сумму в левой и правой части.
		3. Ищем максимальную сумму подмассива, который пересекает границу.
		4. Возвращаем максимум из трех значений: левая часть, правая часть и пересекающая граница.

		ТАЙМ-КОМПЛЕКСНОСТЬ: O(n log n), где n — количество элементов в массиве.
		СПЕЙС-КОМПЛЕКСНОСТЬ: O(log n) для рекурсии.
	*/
	if left == right {
		return arr[left]
	}

	mid := (left + right) / 2

	leftSum := maxSubArraySum(arr, left, mid)
	rightSum := maxSubArraySum(arr, mid+1, right)
	crossSum := maxCrossingSum(arr, left, right, mid)

	return max(leftSum, max(rightSum, crossSum))
}

// Вспомогательная функция для нахождения максимума из двух чисел
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Пример: находим максимальную сумму подмассива
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// Вызываем функцию для нахождения максимальной суммы подмассива
	result := maxSubArraySum(arr, 0, len(arr)-1)

	// Выводим результат
	fmt.Printf("Максимальная сумма подмассива: %d\n", result)
}
