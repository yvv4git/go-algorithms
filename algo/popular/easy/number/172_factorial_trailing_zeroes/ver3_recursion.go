package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func trailingZeroesV4(n int) int {
	/*
		METHOD: Recursion

		TIME COMPLEXITY: O(n), так как мы вызываем функцию factorial(n), которая вычисляет факториал числа n.
		Для каждого числа от n до 1 мы вызываем функцию, поэтому общее количество вызовов функции будет равно n.

		SPACE COMPLEXITY: O(n), так как мы используем рекурсию для вычисления факториала,
		и для каждого рекурсивного вызова требуется свободное место в стеке.
	*/
	fact := factorial(n)
	count := 0
	for fact%10 == 0 {
		count++
		fact /= 10
	}
	return count
}

func main() {
	fmt.Println(trailingZeroes(30)) // Вывод: 7
}
