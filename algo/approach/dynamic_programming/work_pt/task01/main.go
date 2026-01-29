package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	if _, err := fmt.Fscan(in, &n); err != nil {
		return
	}

	result := countRoutes(n)
	fmt.Println(result)
}

func countRoutes(n int) int {
	/*
		METHOD: Dynamic Programming (оптимизированное решение с 3 переменными)
		Суть задачи: Мячик прыгает с N ступенек вниз, может прыгнуть на 1, 2 или 3 ступеньки.
		Нужно найти количество возможных маршрутов от вершины до земли.

		TIME COMPLEXITY: O(n) - один проход по всем ступенькам
		SPACE COMPLEXITY: O(1) - используем только 3 переменные вместо массива
	*/

	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	a, b, c := 1, 1, 2
	for i := 3; i <= n; i++ {
		a, b, c = b, c, a+b+c
	}
	return c
}
