package main

import (
	"fmt"
)

// Функция weirdAlgorithm принимает целое число n и возвращает последовательность чисел,
// полученную применением к n определенных правил.
func weirdAlgorithm(n int) []int {
	// Инициализируем слайс sequence, который будет содержать последовательность чисел.
	// Первым элементом последовательности будет число n.
	sequence := []int{n}

	// Запускаем цикл, который будет выполняться до тех пор, пока n не станет равным 1.
	for n != 1 {
		// Если n четное, то делим его на 2.
		if n%2 == 0 {
			n = n / 2
		} else {
			// Если n нечетное, то умножаем его на 3 и прибавляем 1.
			n = n*3 + 1
		}
		// Добавляем полученное число в конец последовательности.
		sequence = append(sequence, n)
	}

	// Возвращаем полученную последовательность чисел.
	return sequence
}

func main() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)
	sequence := weirdAlgorithm(n)
	fmt.Println("Последовательность чисел:", sequence)
}
