package main

import (
	"bufio"
	"fmt"
	"os"
)

func minBaseLength(t string) int {
	/*
		TASK: Найти минимально возможную длину строки S, из повторений которой получен данный префикс t.
		Строка S писалась несколько раз подряд, после чего от результата взяли префикс t.
		Нужно вернуть минимальную длину S.
		Пример: t = "zzz" - это префикс "zzzz...", поэтому S = "z" и ответ 1.
		Пример: t = "bcabcab" - это префикс "bcabcabc...", поэтому S = "bca" и ответ 3.
		Пример: t = "ababab" - это префикс "abababab...", поэтому S = "ab" и ответ 2.
		Пример: t = "abc" - S может быть "abc", повторённым много раз, поэтому ответ 3.

		METHOD: Bruteforce (перебор периода)
		Полный перебор - это проверка всех возможных вариантов по очереди.
		Перебираем длину p от 1 до n и проверяем, что t[i] = t[i mod p] для всех i.
		Это условие означает, что t является префиксом повторений строки t[:p].
		Первый подходящий p и есть минимальная длина S.

		TIME COMPLEXITY: O(n^2)
		- n возможных длин p, каждая проверка занимает до O(n).

		SPACE COMPLEXITY: O(1)
		- Дополнительные массивы не нужны, только счётчики и флаги.
	*/
	n := len(t)
	if n == 0 {
		return 0
	}

	for p := 1; p <= n; p++ {
		matches := true
		for i := 0; i < n; i++ {
			if t[i] != t[i%p] {
				matches = false
				break
			}
		}
		if matches {
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

	fmt.Fprintln(out, minBaseLength(t))
}
