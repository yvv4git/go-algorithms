package main

import (
	"fmt"
)

func addToArrayForm(num []int, k int) []int {
	/*
		METHOD: Loop
		- Мы используем цикл для сложения каждой цифры массива num с соответствующей цифрой числа k.

		  TIME COMPLEXITY: O(max(len(num), log(k)))
		Временная сложность зависит от максимальной длины между массивом num и количеством цифр в числе k.

		SPACE COMPLEXITY: O(max(len(num), log(k)))
		Пространственная сложность также зависит от максимальной длины между массивом num и количеством цифр в числе k,
		так как мы создаем новый массив для хранения результата.

	*/
	// Инициализация переменной для хранения результата
	result := make([]int, 0)

	// Инициализация переменной для хранения переноса
	carry := 0

	// Индекс для прохода по массиву num с конца
	i := len(num) - 1

	// Пока есть цифры в num или k не стало равным 0 или есть перенос
	for i >= 0 || k > 0 || carry > 0 {
		// Сумма текущей цифры num, цифры k и переноса
		sum := carry
		if i >= 0 {
			sum += num[i]
			i--
		}
		if k > 0 {
			sum += k % 10
			k /= 10
		}

		// Вычисление нового переноса и текущей цифры результата
		carry = sum / 10
		result = append([]int{sum % 10}, result...)
	}

	return result
}

func main() {
	num := []int{1, 2, 0, 0}
	k := 34
	fmt.Println(addToArrayForm(num, k)) // Вывод: [1, 2, 3, 4]

	num = []int{2, 7, 4}
	k = 181
	fmt.Println(addToArrayForm(num, k)) // Вывод: [4, 5, 5]

	num = []int{2, 1, 5}
	k = 806
	fmt.Println(addToArrayForm(num, k)) // Вывод: [1, 0, 2, 1]
}
