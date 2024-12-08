package main

import (
	"fmt"
)

func numPrimeArrangements(n int) int {
	/*
		METHOD: Sieve of Eratosthenes

		Используем решето Эратосфена для подсчета простых чисел до n.
		Затем вычисляем факториалы количества простых и непростых чисел.

		TIME COMPLEXITY: O(n log log n)
		Процесс отметки составных чисел для каждого простого числа `i` требует `n/i` операций.
		Суммарное количество операций для всех простых чисел до `n` составляет `n/2 + n/3 + n/5 + ...`,
		что асимптотически равно `n log log n`.

		SPACE COMPLEXITY: O(n)
		Используется массив `nums` размером `n+1` для отметки составных чисел.
	*/
	data := make([]bool, n+1)

	var c int
	for i := 2; i <= n; i++ {
		if !data[i] {
			c++
			for j := i; j+i <= n; {
				data[j+i] = true
				j += i
			}
		}
	}

	return factorial(c) * factorial(n-c) % (1e9 + 7)
}

func factorial(n int) int {
	/*
		METHOD: Recursion

		Вычисляем факториал числа n рекурсивно.

		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n) (за счет стека вызовов)
	*/
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1) % (1e9 + 7)
}

func main() {
	n := 100 // Пример значения n
	fmt.Printf("Количество допустимых перестановок для первых %d цифр числа π: %d\n", n, numPrimeArrangements(n))
}
