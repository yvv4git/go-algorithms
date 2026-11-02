package main

import (
	"bufio"
	"fmt"
	"os"
)

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	left, right := 0, 0
	for i := 1; i < n; i++ {
		if i <= right {
			z[i] = min(right-i+1, z[i-left])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > right {
			left = i
			right = i + z[i] - 1
		}
	}
	return z
}

func minBaseLengthV3(t string) int {
	/*
		TASK: Найти минимально возможную длину строки S, из повторений которой получен данный префикс t.
		Строка S писалась несколько раз подряд, после чего от результата взяли префикс t.
		Нужно вернуть минимальную длину S.
		Пример: t = "zzz" - это префикс "zzzz...", поэтому S = "z" и ответ 1.
		Пример: t = "bcabcab" - это префикс "bcabcabc...", поэтому S = "bca" и ответ 3.
		Пример: t = "ababab" - это префикс "abababab...", поэтому S = "ab" и ответ 2.
		Пример: t = "abc" - S может быть "abc", повторённым много раз, поэтому ответ 3.

		METHOD: Z-function (Z-функция)
		Z-функция Z[i] - это длина наибольшего общего префикса строки t и её суффикса t[i:].
		Проще говоря, Z[i] показывает, сколько первых символов строки совпадает с символами начиная с i.
		Если Z[p] >= n - p, то t[p:] совпадает с t[:n-p], значит t продолжается с периодом p.
		Перебираем p от 1 до n-1 и возвращаем первое подходящее p.
		Если такого p нет, то весь префикс t и есть минимальная строка S, ответ равен n.

		TIME COMPLEXITY: O(n)
		- Z-функция считается за один проход, внутренний цикл амортизирован.

		SPACE COMPLEXITY: O(n)
		- Массив Z-функции длины n.
	*/
	n := len(t)
	if n == 0 {
		return 0
	}

	z := zFunction(t)
	for p := 1; p < n; p++ {
		if z[p] >= n-p {
			return p
		}
	}
	return n
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t string
	if _, err := fmt.Fscan(in, &t); err != nil {
		return
	}

	fmt.Fprintln(out, minBaseLengthV3(t))
}
