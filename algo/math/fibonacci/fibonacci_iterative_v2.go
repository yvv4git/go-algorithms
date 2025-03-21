//go:build ignore

package main

import (
	"fmt"
)

func fibIterative(n int) int {
	/*
		APPROACH: Iterative

		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(1)
	*/
	if n <= 1 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		next := prev + curr
		prev, curr = curr, next
	}

	return curr
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	fmt.Printf("Число Фибоначчи для %d: %d\n", n, fibIterative(n))
}
