package main

import "fmt"

func ZFunction(s string) []int {
	/*
		METHOD: Z-функция
		TIME COMPLEXITY: O(n)
		SPACE COMPLEXITY: O(n)

		z[i] — длина наибольшего совпадения префикса s и подстроки s[i..n-1].
		z[0] обычно принимается равным n (или 0 — в разных определениях).

		Использует "отрезок совпадений" [l, r): при вычислении z[i] внутри отрезка
		берётся начальное значение из z[i-l], после чего выполняется наивное расширение.
		Благодаря отрезку каждый символ строки сравнивается не более одного раза.
	*/
	n := len(s)
	z := make([]int, n)

	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i < r {
			z[i] = r - i
			if z[i-l] < z[i] {
				z[i] = z[i-l]
			}
		}

		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		if i+z[i] > r {
			l, r = i, i+z[i]
		}
	}

	return z
}

func main() {
	/*
		Пример вычисления Z-функции для строк.
		Z-функция — прямой аналог префикс-функции, выражаются друг через друга.
	*/
	strings := []string{
		"AAAA",
		"ABABAB",
		"ABCDABC",
		"ABACABAB",
	}

	for _, s := range strings {
		z := ZFunction(s)

		fmt.Printf("Строка: %s\n", s)
		fmt.Printf("z:      %v\n\n", z)
	}
}
