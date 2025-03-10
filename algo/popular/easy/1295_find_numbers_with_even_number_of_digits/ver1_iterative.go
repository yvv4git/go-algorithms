//go:build ignore

package main

import (
	"fmt"
)

// Функция findNumbers принимает массив целых чисел nums и возвращает количество чисел с четным количеством цифр.
func findNumbers(nums []int) int {
	/*
		METHOD: Iteration
		1. Перебираем каждое число в массиве.
		2. Для каждого числа считаем количество цифр.
		3. Если количество цифр четное, увеличиваем счетчик.
		4. Возвращаем общее количество чисел с четным количеством цифр.

		TIME COMPLEXITY: O(n * k), где n — количество элементов в массиве, k — среднее количество цифр в числах.

		SPACE COMPLEXITY: O(1), так как используется фиксированное количество дополнительной памяти.
	*/
	count := 0 // Переменная для подсчета чисел с четным количеством цифр

	// Перебираем каждое число в массиве
	for _, num := range nums {
		// Вычисляем количество цифр в текущем числе
		digits := countDigits(num)

		// Если количество цифр четное, увеличиваем счетчик
		if digits%2 == 0 {
			count++
		}
	}

	return count // Возвращаем итоговое количество чисел с четным количеством цифр
}

// Вспомогательная функция для подсчета количества цифр в числе
func countDigits(num int) int {
	/*
		METHOD: Mathematical calculation
		TIME COMPLEXITY: O(k), где k — количество цифр в числе.
		SPACE COMPLEXITY: O(1), так как используется фиксированное количество дополнительной памяти.
	*/
	digits := 0

	// Пока число больше нуля, делим его на 10 и увеличиваем счетчик цифр
	for num > 0 {
		num /= 10
		digits++
	}

	// Обрабатываем случай, когда число равно 0 (хотя по условию задачи это невозможно)
	if digits == 0 {
		return 1
	}

	return digits
}

func main() {
	// Пример входных данных
	nums := []int{12, 345, 2, 6, 7896}

	// Вызываем функцию и выводим результат
	result := findNumbers(nums)
	fmt.Println(result) // Ожидаемый вывод: 2
}
