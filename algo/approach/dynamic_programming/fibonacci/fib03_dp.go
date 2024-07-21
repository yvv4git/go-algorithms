package main

import (
	"fmt"
)

// fibV3 вычисляет n-ое число Фибоначчи с использованием динамического программирования.
// Временная сложность: O(n)
// Пространственная сложность: O(1)
func fibV3(n int) int {
	a := 1 // f(i - 2)
	b := 1 // f(i - 1)
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println("Fibonacci result:", fibV3(5))
}
