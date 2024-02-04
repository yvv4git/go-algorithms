package main

import (
	"unicode"
)

func calculate(s string) int {
	/*
		METHOD: Stack
		Time complexity: O(n)
		Space complexity: O(n)

	*/
	// Стек для хранения чисел и знаков операций
	stack := make([]int, 0)
	num := 0
	sign := '+'

	for i, r := range s {
		// Если символ - цифра, то формируем число
		if unicode.IsDigit(r) {
			num = num*10 + int(r-'0')
		}

		// Если символ не цифра и не пробел, или это конец строки,
		// то нужно выполнить действие с числом и обновить знак
		if (!unicode.IsDigit(r) && !unicode.IsSpace(r)) || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				// Умножаем последнее число в стеке на текущее число
				stack[len(stack)-1] *= num
			case '/':
				// Делим последнее число в стеке на текущее число
				stack[len(stack)-1] /= num
			}

			// Обновляем знак и сбрасываем число
			sign = r
			num = 0
		}
	}

	// Вычисляем сумму всех чисел в стеке
	result := 0
	for _, val := range stack {
		result += val
	}

	return result
}

func main() {
	s := "3+2*2"
	result := calculate(s)
	println(result) // Должно вывести 7
}
