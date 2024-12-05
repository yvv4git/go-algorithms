package main

import (
	"fmt"
	"strconv"
)

// Функция для проверки, содержит ли число цифру 0
func hasZero(num int) bool {
	// Преобразуем число в строку и проверяем, есть ли в ней символ '0'
	strNum := strconv.Itoa(num)
	for _, char := range strNum {
		if char == '0' {
			return true
		}
	}
	return false
}

// Основная функция, которая находит два числа без нулей, сумма которых равна n
func getNoZeroIntegers(n int) []int {
	/*
		METHOD: Linear Search with String Conversion
		- Мы используем линейный поиск, начиная с чисел a = 1 и b = n - 1.
		- Для каждой пары чисел a и b проверяем, содержат ли они цифру 0.
		- Если хотя бы одно из чисел содержит цифру 0, увеличиваем a на 1 и уменьшаем b на 1, пока не найдем пару чисел без нулей.
		- Этот метод основан на преобразовании чисел в строки и посимвольной проверке на наличие нуля.

		TIME COMPLEXITY: O(n)
		- В худшем случае, мы проверим все числа от 1 до n/2, что дает нам временную сложность O(n).

		SPACE COMPLEXITY: O(1)
		- Пространственная сложность также O(1),
		так как мы используем фиксированное количество дополнительной памяти (переменные a, b
		и временная строка для проверки на ноль).
	*/
	// Начинаем с чисел a = 1 и b = n - 1
	a := 1
	b := n - 1

	// Пока одно из чисел содержит цифру 0, продолжаем искать
	for hasZero(a) || hasZero(b) {
		// Увеличиваем a на 1 и уменьшаем b на 1
		a++
		b--
	}

	// Возвращаем найденные числа a и b
	return []int{a, b}
}

func main() {
	// Пример использования функции
	n := 11
	result := getNoZeroIntegers(n)
	fmt.Println(result) // Вывод: [2, 9]
}
