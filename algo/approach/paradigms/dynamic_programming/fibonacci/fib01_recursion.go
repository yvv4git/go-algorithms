package main

import (
	"errors"
)

// Fib вычисляет n-ое число Фибоначчи с использованием рекурсии.
// Временная сложность: O(2^n)
// Пространственная сложность: O(n)
func fib(n int) (int, error) {
	// Проверка на отрицательное число
	if n < 0 {
		return 0, errors.New("negative number is not allowed")
	}

	// Базовые случаи
	if n == 0 {
		return 1, nil
	}

	if n == 1 {
		return 1, nil
	}

	// Рекурсивные вызовы
	fib1, _ := fib(n - 1)
	fib2, _ := fib(n - 2)

	return fib1 + fib2, nil
}

//func main() {
//	result, err := fib(5)
//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("Fibonacci result:", result)
//	}
//}
