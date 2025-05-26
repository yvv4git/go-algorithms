//go:build ignore

package main

import (
	"fmt"
)

// Хвостовая рекурсия (tail recursion) — это особый случай рекурсии,
// при котором рекурсивный вызов является последней операцией в функции.
//
// Это нужно для того:
// - В некоторых языках программирования компилятор может автоматически преобразовать хвостовую рекурсию в цикл,
// тем самым экономя стек памяти и предотвращая переполнение стека (stack overflow).
//
// В данном примере, acc - это аккумулятор, который накапливает результат.

func Factorial(n int) int {
	return factorialTailRec(n, 1)
}

func factorialTailRec(n, acc int) int {
	if n <= 1 {
		return acc // Базовый случай(база рекурсии)
	}

	return factorialTailRec(n-1, acc*n)
}

func main() {
	n := 5
	result := Factorial(n)
	fmt.Printf("Факториал числа %d равен %d\n", n, result)
}
